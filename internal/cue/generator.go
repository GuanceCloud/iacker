package cue

import (
	"errors"
	"fmt"
	"path/filepath"

	"cuelang.org/go/cue"
	"github.com/fatih/color"
	"github.com/hashicorp/go-multierror"

	"github.com/GuanceCloud/iacker/internal/helpers/filediff"
	templateV1 "github.com/GuanceCloud/iacker/pkg/template/v1"
)

type generator struct{}

// GenerateOption is a function that configures a generator.
type GenerateOption func(*generator) error

// Generate will generate code files into the output folder
func Generate(p *Package, options ...GenerateOption) error {
	plan, err := GeneratePlan(p, options...)
	if err != nil {
		return fmt.Errorf("failed to generate plan: %w", err)
	}
	if err := plan.Pretty(); err != nil {
		return fmt.Errorf("failed to pretty print plan: %w", err)
	}
	return plan.Save()
}

// GeneratePlan will generate a plan for generating code files into the output folder
func GeneratePlan(p *Package, options ...GenerateOption) (*Plan, error) {
	var mErr error
	sourceFiles := make(map[string]string)
	targetFiles := make(map[string]string)
	rms := p.Manifest()

	// Create generator
	g := &generator{}
	for _, option := range options {
		if err := option(g); err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("generate option failed: %w", err))
		}
	}

	// Generate files from templates
	for i, templateOptions := range rms.Options.Templates {
		templateName := templateOptions.Template

		color.Magenta("Generate template %q to folder %q", templateName, templateOptions.Outdir)

		// Get the template
		templateValue := p.value.LookupPath(cue.MakePath(cue.Str("templates"), cue.Str(templateName)))
		if err := templateValue.Err(); err != nil {
			return nil, fmt.Errorf("lookup path failed: %w", err)
		}

		// Fill the template with the input RMS data
		concretTemplateValue := templateValue.FillPath(cue.MakePath(cue.Str("inputs")), &templateV1.Inputs{
			Resources: rms.Resources,
			Errors:    rms.Errors,
			Vars:      templateOptions.Vars,
		})
		if err := concretTemplateValue.Err(); err != nil {
			return nil, fmt.Errorf("fill path failed: %w", err)
		}

		// Check the template diagnostics
		var concretTemplate templateV1.Manifest
		if err := concretTemplateValue.Decode(&concretTemplate); err != nil {
			return nil, fmt.Errorf("decode template %s(%d): %w", templateName, i, err)
		}
		for _, diag := range concretTemplate.Diagnostics {
			mErr = multierror.Append(mErr, errors.New(diag.Message))
		}

		// Build the generate plan
		files, err := filediff.ReadFiles(templateOptions.Outdir)
		if err != nil {
			return nil, fmt.Errorf("read existed files from outdir failed: %w", err)
		}
		for fileName, file := range files {
			sourceFiles[filepath.Join(templateOptions.Outdir, fileName)] = file
		}

		for fileName, file := range concretTemplate.Outputs.Files {
			targetFiles[filepath.Join(templateOptions.Outdir, fileName)] = file.Content
		}
	}
	if mErr != nil {
		return nil, mErr
	}

	return &Plan{sources: sourceFiles, targets: targetFiles}, nil
}
