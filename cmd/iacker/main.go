package main

import (
	"fmt"

	"github.com/GuanceCloud/iacker/internal/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
