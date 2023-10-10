package cue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePlan(t *testing.T) {
	rms, err := ParsePackage("../../examples/simple")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	plan, err := GeneratePlan(rms)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	if !assert.NotNil(t, plan) {
		t.Fatal("plan is nil")
	}

	plan.Pretty()
}

func TestGenerate(t *testing.T) {
	rms, err := ParsePackage("../../examples/simple")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	err = Generate(rms)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}
	assert.FileExists(t, ".build/README.md")
}
