package stringutils

import "strings"

// RemoveWhitespace removes all whitespace from a string.
func RemoveWhitespace(s string) string {
	for _, r := range []rune{'\t', '\r', '\n', ' '} {
		s = strings.ReplaceAll(s, string(r), "")
	}
	return s
}
