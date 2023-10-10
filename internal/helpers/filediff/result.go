package filediff

import (
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// Result represents the diff result.
type Result struct {
	Created   []string
	Deleted   []string
	Unchanged []string
	Diffs     map[string]Diff
}

// HasChanged returns true if the result has any changes.
func (r Result) HasChanged() bool {
	return len(r.Created) > 0 || len(r.Diffs) > 0
}

// Pretty will print the diff result.
func (r Result) Pretty(w io.Writer) {
	color.Output = w
	for _, path := range r.Created {
		color.Green("Created: %s", path)
	}

	for range r.Deleted {
		// color.Red("Deleted: %s", path)
	}

	for _, path := range r.Unchanged {
		color.Cyan("Unchanged: %s", path)
	}

	for path, d := range r.Diffs {
		color.Red("--- %s\n", path)
		color.Green("+++ %s\n", path)
		d.Pretty(w)
	}
}

// Diff represents the diff between two files.
type Diff struct {
	Old string
	New string
}

// Pretty will print the diff.
func (diff Diff) Pretty(w io.Writer) {
	color.Output = w

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(diff.Old, diff.New, true)
	fmt.Println(dmp.DiffPrettyText(diffs))
}
