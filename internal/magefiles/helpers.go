//go:build mage
// +build mage

package main

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/go-multierror"
	"github.com/magefile/mage/sh"
)

func batchRun(commands []string) error {
	var mErr error
	for _, command := range commands {
		program, args := parseCommand(command)
		if err := sh.Run(program, args...); err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("run %q failed: %w", command, err))
		}
	}
	return mErr
}

func parseCommand(command string) (string, []string) {
	parts := regexp.MustCompile(`\s+`).Split(command, -1)
	return parts[0], parts[1:]
}
