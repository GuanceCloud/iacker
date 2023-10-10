package filediff

import (
	"os"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
)

type Files map[string]string

// ReadFiles reads all files under the specified directory.
func ReadFiles(root string) (Files, error) {
	var mErr error
	files := Files{}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			mErr = multierror.Append(mErr, err)
			return nil
		}
		if info.IsDir() {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			mErr = multierror.Append(mErr, err)
			return nil
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			mErr = multierror.Append(mErr, err)
			return nil
		}
		files[relPath] = string(content)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// Save saves the outputs to the specified directory.
func (files Files) Save() error {
	var mErr error
	for path, fileContent := range files {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			mErr = multierror.Append(mErr, err)
		}

		if err := os.WriteFile(path, []byte(fileContent), 0o600); err != nil {
			mErr = multierror.Append(mErr, err)
		}
	}
	return mErr
}
