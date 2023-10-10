package cue

import (
	"cuelang.org/go/cue"
	rmsV1 "github.com/GuanceCloud/iacker/pkg/rms/v1"
)

// Package is a package written by CUE.
type Package struct {
	root     string
	value    *cue.Value
	manifest *rmsV1.Manifest
}

// Manifest returns the resource management specification of the package.
func (pkg *Package) Manifest() *rmsV1.Manifest {
	return pkg.manifest
}
