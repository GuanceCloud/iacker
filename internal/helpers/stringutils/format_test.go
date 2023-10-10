package stringutils

import "testing"

func TestRemoveWhitespace(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "removes spaces",
			input:    "hello world",
			expected: "helloworld",
		},
		{
			name:     "removes tabs",
			input:    "hello\tworld",
			expected: "helloworld",
		},
		{
			name:     "removes newlines",
			input:    "hello\nworld",
			expected: "helloworld",
		},
		{
			name:     "removes carriage returns",
			input:    "hello\rworld",
			expected: "helloworld",
		},
		{
			name:     "removes all whitespace",
			input:    "  \t\n\rhello \t\n\rworld  \t\n\r",
			expected: "helloworld",
		},
		{
			name:     "handles empty string",
			input:    "",
			expected: "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := RemoveWhitespace(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %q but got %q", tc.expected, actual)
			}
		})
	}
}
