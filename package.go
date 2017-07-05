package goref

import (
	"go/token"
	"go/types"
	"golang.org/x/tools/go/loader"
)

// Package represents a Go Package, including its dependencies.
type Package struct {
	// Name of the package
	Name string

	// Dependencies is the map of load-paths imported within this
	// package to the corresponding Package objects.
	Dependencies map[string]*Package

	// Dependents is the map of packages' load-paths that load
	// this package through any load-path, mapped to their
	// corresponding Package objects.
	Dependents map[string]*Package

	// Files is a map of paths to File objects that make up this package.
	Files map[string]*File

	// OutRefs and InRefs are maps of references from a Position
	// (file:line:column). For OutRefs the Position is local to
	// the package and the Ref is to an identifier in another
	// package. For InRefs the Position is external to the package
	// and the Ref is to an identifier within this package.
	OutRefs map[Position]*Ref
	InRefs  map[Position]*Ref

	// Interfaces is the list of interface types in this package.
	//
	// This is used to compute the interface-implementation matrix.
	//
	// Only named interfaces matter, because an unnamed interface
	// can't be exported.
	//
	// Interfaces equivalent to interface{} are excluded.
	Interfaces []*types.Named

	// Impls is the list of non-interface types in this package.
	//
	// This is used to compute the interface-implementation matrix.
	//
	// Only named types matter, because an unnamed type can't have
	// methods.
	Impls []*types.Named

	// Fset is a reference to the token.FileSet that loaded this
	// package.
	Fset *token.FileSet
}

func (p *Package) String() string {
	return p.Name
}

func newPackage(pi *loader.PackageInfo, fset *token.FileSet) *Package {
	return &Package{
		//PackageInfo:  pi,
		Name:         pi.Pkg.Name(),
		Dependencies: make(map[string]*Package),
		Dependents:   make(map[string]*Package),
		Files:        make(map[string]*File),
		OutRefs:      make(map[Position]*Ref),
		InRefs:       make(map[Position]*Ref),
		Interfaces:   make([]*types.Named, 0),
		Impls:        make([]*types.Named, 0),
		Fset:         fset,
	}
}