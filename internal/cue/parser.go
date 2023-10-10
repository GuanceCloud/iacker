package cue

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/errors"
	"cuelang.org/go/cue/load"
	rmsV1 "github.com/GuanceCloud/iacker/pkg/rms/v1"
)

type parser struct{}

// ParseOption is a function that configures a parser.
type ParseOption func(*parser) error

// ParsePackage parses the package.
func ParsePackage(root string, options ...ParseOption) (*Package, error) {
	p := &parser{}
	for _, option := range options {
		if err := option(p); err != nil {
			return nil, fmt.Errorf("apply option failed: %w", err)
		}
	}

	v, err := p.parseCUE(root)
	if err != nil {
		return nil, fmt.Errorf("parse CUE package failed: %w", err)
	}

	rms := &rmsV1.Manifest{}
	if err := v.Decode(rms); err != nil {
		return nil, fmt.Errorf("decode CUE package into manifest failed, %w", err)
	}

	return &Package{
		root:     root,
		value:    v,
		manifest: rms,
	}, nil
}

func (p *parser) parseCUE(path string) (*cue.Value, error) {
	// We need a cue.Context for building after loading
	ctx := cuecontext.New()

	// The entrypoints are the same as the files you'd specify at the command line
	entrypoints := []string{}
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".cue") {
			continue
		}
		entrypoints = append(entrypoints, filepath.Join(path, file.Name()))
	}

	// Load Cue files into Cue build.Instances slice
	// the second arg is a configuration object, we'll see this later
	bis := load.Instances(entrypoints, nil)

	values, err := ctx.BuildInstances(bis)
	if err != nil {
		cueErr, ok := err.(errors.Error)
		if ok {
			return nil, fmt.Errorf("build instances failed, %w, %s", err, cueErr.Position())
		}
		return nil, fmt.Errorf("build instances failed, %w", err)
	}

	if len(values) > 1 {
		return nil, fmt.Errorf("more than one value found")
	}

	value := values[0]

	// Validate the value
	err = value.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate value failed, %w", err)
	}
	return &value, nil
}
