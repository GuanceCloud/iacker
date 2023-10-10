package cue

import (
	"fmt"
	"os"

	"github.com/GuanceCloud/iacker/internal/helpers/filediff"
)

type Plan struct {
	sources map[string]string
	targets map[string]string
}

// Save saves the plan to the output folder
func (plan *Plan) Save() error {
	return filediff.Files(plan.targets).Save()
}

// Pretty prints the plan to stdout
func (plan *Plan) Pretty() error {
	// Diff between output and snapshot
	result, err := filediff.FileDiff(plan.sources, plan.targets)
	if err != nil {
		return fmt.Errorf("diff error: %w", err)
	}
	result.Pretty(os.Stdout)
	return nil
}
