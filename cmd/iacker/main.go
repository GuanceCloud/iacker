package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "iacker",
		Short: "The IaC Provider Development Framework",
		Long:  `Help you to develop IaC provider easily.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, world!")
		},
	}
}
