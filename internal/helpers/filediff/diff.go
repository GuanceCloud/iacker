package filediff

import "github.com/GuanceCloud/iacker/internal/helpers/stringutils"

// FileDiff compare the actual outputs with the snapshot.
func FileDiff(sources Files, targets Files) (*Result, error) {
	result := &Result{
		Diffs: make(map[string]Diff),
	}

	// diff two maps, get the created, updated and deleted files.
	for path, otherFile := range targets {
		snapshotFile, ok := sources[path]
		if !ok {
			result.Created = append(result.Created, path)
			continue
		}

		if stringutils.RemoveWhitespace(otherFile) != stringutils.RemoveWhitespace(snapshotFile) {
			result.Diffs[path] = Diff{
				New: otherFile,
				Old: snapshotFile,
			}
		} else {
			result.Unchanged = append(result.Unchanged, path)
		}
	}

	for path := range sources {
		if _, ok := targets[path]; !ok {
			result.Deleted = append(result.Deleted, path)
		}
	}
	return result, nil
}
