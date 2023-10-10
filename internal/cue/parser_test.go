package cue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePackage(t *testing.T) {
	tests := []struct {
		name    string
		root    string
		options []ParseOption
		wantErr bool
	}{
		{
			name: "parse package",
			root: "../../examples/simple",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePackage(tt.root, tt.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePackage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotEmpty(t, got, "package is empty")
		})
	}
}
