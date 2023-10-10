package cmd

import (
	"github.com/spf13/cobra"
)

// NewRootCmd returns a new root command.
func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "iacker",
		Short: "The IaC Provider Development Framework",
		Long:  `Help you to develop IaC provider easily.`,
	}
	root.AddCommand(
		NewInitCmd(),
		NewGenerateCmd(),
	)
	return root
}
