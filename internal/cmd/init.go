package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewInitCmd returns a new init command.
func NewInitCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new IaC provider project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("not implemented")
		},
	}
	return initCmd
}
