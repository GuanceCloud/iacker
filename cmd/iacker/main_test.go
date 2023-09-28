package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	rootCmd := NewRootCmd()
	rootCmd.SetOutput(&buf)
	rootCmd.SetArgs([]string{"--help"})
	if err := rootCmd.Execute(); err != nil {
		t.Errorf("error executing command: %v", err)
	}
	output := buf.String()
	if output == "" {
		t.Errorf("expected output, but got none")
	}
}
