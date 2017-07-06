package goref

import (
	log "github.com/sirupsen/logrus"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/loader"
	"path"
	"strings"
)

// PackageGraph represents a collection of Go packages and their
// mutual dependencies. All dependencies of a Package in the
// PackageGraph are also part of the PackageGraph.
type PackageGraph struct {
	// Map of package load-path to Package objects.
	Packages map[string]*Package

	// Map of file path to File objects.
	Files map[string]*File
}

// cleanImportSpec takes an ast.ImportSpec and cleans the Path
// component by trimming the quotes (") that surround it.
func cleanImportSpec(spec *ast.ImportSpec) string {
	// FIXME we should make sure we understand what can cause Path
	// to be empty.
	if spec.Path != nil {
		s := spec.Path.Value
		s = strings.Trim(s, "\"")
		return s
	}
	return "<unknown>"
}

// candidatePaths returns a slice enumerating all the possible import
// paths for a package. This means inserting the possible "vendor"
// directory location from the load path of the importing package.
//
// If package a/b imports c/d, the following paths are candidates:
// a/b/vendor/c/d
// a/vendor/c/d
// vendor/c/d
// c/d
//
// Order matters, as the most-specific vendored package is selected.
// Note that multi-level vendoring works, as PackageGraph will
// consider the full import path, including path/to/vendor/, as the
// package path when building the graph. In this sense we follow the
// go tool's convention to not try to detect when two packages loaded
// through different paths are the same package.
func candidatePaths(loadpath, parent string) []string {
	const kVendor = "vendor"
	paths := []string{}
	for parent != "." && parent != "" {
		paths = append(paths, path.Join(parent, kVendor, loadpath))
		parent = path.Dir(parent)
	}
	// Some dependencies may be vendored under
	// $GOROOT/src/vendor. This is the case e.g. for
	// `golang_org/x/net/lex/httplex` which is imported by
	// `net/http`. This is the correct path: it's just vendored
	// that way in the standard library. See
	// https://github.com/golang/go/issues/16333 for background on
	// that.
	paths = append(paths, path.Join(kVendor, loadpath))
	paths = append(paths, loadpath)
	return paths
}

// loadPackage recursively loads a Go package into the Package
// Graph. If the package was already loaded, it returns early. It
// always returns the Package object for the loaded package.
func (pg *PackageGraph) loadPackage(prog *loader.Program, loadpath string, pi *loader.PackageInfo) *Package {
	if pkg, in := pg.Packages[loadpath]; in {
		return pkg
	}
	pkg := newPackage(pi, prog.Fset)
	pg.Packages[loadpath] = pkg

	// Iterate over all files in that package.
	for _, f := range pi.Files {
		// Get the File object to represent that file
		filepath := prog.Fset.File(f.Package).Name()
		ff := newFile(prog.Fset)

		// Add this File to the maps for the Package and for
		// the PackageGraph.
		pkg.Files[filepath] = ff
		pg.Files[filepath] = ff

		// Iterate over all imports in that file
		for _, imported := range f.Imports {
			// Find the import's load-path and load that
			// package into the graph.
			ipath := cleanImportSpec(imported)
			candidatePaths := candidatePaths(ipath, loadpath)
			var i *loader.PackageInfo
			for _, c := range candidatePaths {
				i = prog.Package(c)
				if i != nil {
					ipath = c
					break
				}
			}
			if i == nil {
				log.Warnf("Tried to load package `%s` imported by package `%s` but it wasn't found anywhere in the load path. The candidate load paths were: %s\n", ipath, loadpath, candidatePaths)
				continue
			}
			importedPkg := pg.loadPackage(prog, ipath, i)

			// Set up the edges on the package dependency graph
			importedPkg.Dependents[loadpath] = pkg
			pkg.Dependencies[ipath] = importedPkg
		}
	}

	// Iterate over all object uses in that package and filter for
	// non-local references only.
	for id, obj := range pi.Uses {
		// the object's Pkg will be nil for builtins
		if obj.Pkg() != nil {
			pkgLoadPath := obj.Pkg().Path()
			if pkgLoadPath != loadpath {
				foreignPkg := pg.Packages[pkgLoadPath]
				if foreignPkg != nil {
					ref := &Ref{
						RefType:      refTypeForId(prog, id),
						ToIdent:      obj.Name(),
						ToPackage:    foreignPkg,
						ToPosition:   NewPosition(prog.Fset, obj.Pos(), NoPos),
						FromIdent:    id.Name,
						FromPackage:  pkg,
						FromPosition: NewPosition(prog.Fset, id.Pos(), id.End()),
					}

					refpos := NewPosition(prog.Fset, id.Pos(), id.End())
					foreignPkg.InRefs[refpos] = ref
					pkg.OutRefs[refpos] = ref
				}
			}
		}
	}

	// Iterate over all types in that package and insert them as
	// needed into Structs and Interfaces.
	for _, name := range pi.Pkg.Scope().Names() {
		if obj, ok := pi.Pkg.Scope().Lookup(name).(*types.TypeName); ok {
			if named, ok := obj.Type().(*types.Named); ok {
				if types.IsInterface(named) {
					i := named.Obj().Type().Underlying().(*types.Interface)
					// We only care about interfaces that are exported, and that are not interface{}.
					if named.Obj().Exported() && i.NumMethods() > 0 {
						pkg.Interfaces = append(pkg.Interfaces, named)
					}
				} else {
					pkg.Impls = append(pkg.Impls, named)
				}
			}
		}
	}

	return pkg
}

// LoadProgram loads recursively packages used from a `main` package.
// It may be called multiple times to load multiple programs'
// package sets in the PackageGraph.
func (p *PackageGraph) LoadProgram(loadpath string, filenames []string) {
	conf := loader.Config{}
	conf.CreateFromFilenames(loadpath, filenames...)

	prog, err := conf.Load()
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range prog.AllPackages {
		p.loadPackage(prog, k.Path(), v)
	}
}

func (p *PackageGraph) ComputeInterfaceImplementationMatrix() {
	for _, pa := range p.Packages {
		for _, iface := range pa.Interfaces {
			for _, pb := range p.Packages {
				for _, typ := range pb.Impls {
					if typ == iface {
						continue
					}
					if types.AssignableTo(typ, iface) {
						fset := pb.Fset
						pos := NewPosition(fset, typ.Obj().Pos(), token.NoPos)
						r := &Ref{
							RefType:      Implementation,
							ToIdent:      iface.Obj().Name(),
							ToPackage:    pa,
							ToPosition:   NewPosition(fset, iface.Obj().Pos(), NoPos),
							FromIdent:    typ.Obj().Name(),
							FromPackage:  pb,
							FromPosition: NewPosition(fset, typ.Obj().Pos(), NoPos),
						}
						pa.InRefs[pos] = r
						pb.OutRefs[pos] = r
					}
				}
				for _, ifaceb := range pb.Interfaces {
					if ifaceb == iface {
						continue
					}
					if types.AssignableTo(ifaceb, iface) {
						fset := pb.Fset
						pos := NewPosition(fset, ifaceb.Obj().Pos(), token.NoPos)
						r := &Ref{
							RefType:    Extension,
							ToIdent:    iface.Obj().Name(),
							ToPackage:  pa,
							ToPosition: NewPosition(fset, ifaceb.Obj().Pos(), NoPos),

							FromIdent:    ifaceb.Obj().Name(),
							FromPackage:  pb,
							FromPosition: NewPosition(fset, ifaceb.Obj().Pos(), NoPos),
						}
						pa.InRefs[pos] = r
						pb.OutRefs[pos] = r
					}
				}
			}
		}
	}
}

// NewPackageGraph returns a new, empty PackageGraph.
func NewPackageGraph() *PackageGraph {
	return &PackageGraph{
		Packages: make(map[string]*Package),
		Files:    make(map[string]*File),
	}
}
