package filediff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFiles(t *testing.T) {
	root := "testdata"
	expectedFiles := Files{
		"test.json":          "",
		"test.snapshot.json": "",
	}

	files, err := ReadFiles(root)
	if err != nil {
		t.Fatalf("ReadFiles failed: %v", err)
	}

	for fileName, content := range files {
		if _, ok := expectedFiles[fileName]; assert.True(t, ok) {
			continue
		}

		if !assert.NotEmpty(t, content) {
			continue
		}
	}
}
