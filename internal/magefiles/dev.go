//go:build mage
// +build mage

package main

import (
	"fmt"
	"path/filepath"

	"github.com/GuanceCloud/iacker/internal/helpers/osfs"
	"github.com/hashicorp/go-multierror"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Dev generate code from RMS definitions
type Dev mg.Namespace

// Fmt format the code
func (dev Dev) Fmt() error {
	commands := []string{
		"gofumpt -l -w .",
		"cue fmt ./...",
	}
	return batchRun(commands)
}

// Lint lint the code
func (dev Dev) Lint() error {
	return sh.Run("golangci-lint", "run", "--fix", "--allow-parallel-runners")
}

// D2 build svg from d2 files
func (dev Dev) D2() error {
	files, err := osfs.ListFileByExt("proposals", "d2")
	if err != nil {
		return err
	}
	var mErr error
	for _, d2File := range files {
		outFile := d2File[:len(d2File)-3] + ".svg"
		if err := sh.RunV("d2", "--sketch", "-t", "0", d2File, outFile); err != nil {
			mErr = fmt.Errorf("d2 svg %s: %w", d2File, err)
		}
	}
	return mErr
}

// Install install the binary into local environment
func (dev Dev) Install() error {
	// Get GOPATH
	gopath, err := sh.Output("go", "env", "GOPATH")
	if err != nil {
		return fmt.Errorf("get GOPATH failed: %w", err)
	}
	binaryName := "iacker"
	return sh.Run(
		"go", "build", "-o",
		filepath.Join(gopath, "bin", binaryName),
		fmt.Sprintf("./cmd/%s", binaryName),
	)
}

// GenCue generate cue from proto files
func (dev Dev) GenCue() error {
	pkgs := []string{
		"rms",
		"resource",
		"template",
	}
	var mErr error
	for _, pkg := range pkgs {
		if err := sh.Run(
			"cue", "import", "-f", "-I", ".",
			fmt.Sprintf("./pkg/%s/v1/%s.proto", pkg, pkg),
		); err != nil {
			mErr = multierror.Append(mErr, err)
		}
	}
	return mErr
}
