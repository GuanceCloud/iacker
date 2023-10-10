package osfs

import (
	"reflect"
	"sort"
	"testing"
)

func TestListFileByExt(t *testing.T) {
	root := "testdata"
	ext := ".txt"
	expectedFiles := []string{"testdata/foo.txt", "testdata/foo/bar.txt"}

	files, err := ListFileByExt(root, ext)
	if err != nil {
		t.Fatalf("ListFileByExt failed: %v", err)
	}

	sort.Strings(files)
	sort.Strings(expectedFiles)

	if !reflect.DeepEqual(files, expectedFiles) {
		t.Fatalf("Expected files %v, but got %v", expectedFiles, files)
	}
}
