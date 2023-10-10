package filediff

import (
	"testing"
)

func TestFileDiff(t *testing.T) {
	files := Files{
		"file1.txt": "hello world",
		"file2.txt": "foo bar",
		"file3.txt": "baz",
	}

	others := Files{
		"file1.txt": "hello world",
		"file2.txt": "foo baz",
		"file4.txt": "new file",
	}

	expectedResult := &Result{
		Diffs: map[string]Diff{
			"file2.txt": {
				New: "foo baz",
				Old: "foo bar",
			},
		},
		Created:   []string{"file4.txt"},
		Deleted:   []string{"file3.txt"},
		Unchanged: []string{"file1.txt"},
	}

	result, err := FileDiff(files, others)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(result.Diffs) != len(expectedResult.Diffs) {
		t.Errorf("expected %d diffs, but got %d", len(expectedResult.Diffs), len(result.Diffs))
	}

	for path, diff := range expectedResult.Diffs {
		if result.Diffs[path].New != diff.New {
			t.Errorf("expected new content of %s to be %s, but got %s", path, diff.New, result.Diffs[path].New)
		}

		if result.Diffs[path].Old != diff.Old {
			t.Errorf("expected old content of %s to be %s, but got %s", path, diff.Old, result.Diffs[path].Old)
		}
	}

	if len(result.Created) != len(expectedResult.Created) {
		t.Errorf("expected %d created files, but got %d", len(expectedResult.Created), len(result.Created))
	}

	for i, path := range expectedResult.Created {
		if result.Created[i] != path {
			t.Errorf("expected created file at index %d to be %s, but got %s", i, path, result.Created[i])
		}
	}

	if len(result.Deleted) != len(expectedResult.Deleted) {
		t.Errorf("expected %d deleted files, but got %d", len(expectedResult.Deleted), len(result.Deleted))
	}

	for i, path := range expectedResult.Deleted {
		if result.Deleted[i] != path {
			t.Errorf("expected deleted file at index %d to be %s, but got %s", i, path, result.Deleted[i])
		}
	}

	if len(result.Unchanged) != len(expectedResult.Unchanged) {
		t.Errorf("expected %d unchanged files, but got %d", len(expectedResult.Unchanged), len(result.Unchanged))
	}

	for i, path := range expectedResult.Unchanged {
		if result.Unchanged[i] != path {
			t.Errorf("expected unchanged file at index %d to be %s, but got %s", i, path, result.Unchanged[i])
		}
	}
}
