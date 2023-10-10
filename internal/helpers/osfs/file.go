package osfs

import (
	"os"
	"path/filepath"
	"strings"
)

// ListFileByExt will list all Cue packages in a directory
func ListFileByExt(root string, ext string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), ext) {
			files = append(files, path)
			return nil
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
