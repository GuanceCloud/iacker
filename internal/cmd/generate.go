package cmd

import (
	"fmt"

	"github.com/GuanceCloud/iacker/internal/cue"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cobra"
)

// GenerateOptions is the options of generate command
type GenerateOptions struct {
	Init bool
}

// NewGenerateCmd returns a new generate command.
func NewGenerateCmd() *cobra.Command {
	options := &GenerateOptions{}
	cmd := &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen"},
		Short:   "Generate code from CUE",
		RunE: func(cmd *cobra.Command, args []string) error {
			var mErr error
			if len(args) == 0 {
				args = []string{"."}
			}

			// Generate all packages.
			for _, root := range args {
				rms, err := cue.ParsePackage(root)
				if err != nil {
					return err
				}
				if err := cue.Generate(
					rms,
					cue.WithLayout(options.Init),
				); err != nil {
					mErr = multierror.Append(mErr, fmt.Errorf("generate %s error: %w", root, err))
				}
			}
			return mErr
		},
	}
	cmd.Flags().BoolVar(&options.Init, "init", false, "initialize project with the layout")
	return cmd
}
