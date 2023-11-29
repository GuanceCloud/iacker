package schema

import (
	"encoding/json"
	"strings"

	"github.com/tidwall/gjson"
)

// JSONPointer is a JSON pointer format described in RFC 6901.
type JSONPointer struct {
	paths []string
}

// NewJSONPointer creates a new JSON pointer.
func NewJSONPointer(paths ...string) *JSONPointer {
	return &JSONPointer{paths: paths}
}

func NewJSONPointerFromRef(ref string) *JSONPointer {
	return NewJSONPointer(strings.Split(strings.TrimPrefix(ref, "/"), "/")...)
}

// Get gets the value of a JSON pointer.
func (p *JSONPointer) Get(content string) string {
	return gjson.Get(content, strings.Join(p.paths, ".")).String()
}

// String returns the string representation of a JSON pointer.
func (p *JSONPointer) String() string {
	return "/" + strings.Join(p.paths, "/")
}

// UnmarshalJSON unmarshal a JSON pointer.
func (p *JSONPointer) UnmarshalJSON(data []byte) error {
	var ref string
	if err := json.Unmarshal(data, &ref); err != nil {
		return err
	}
	p.paths = strings.Split(strings.TrimPrefix(ref, "/"), "/")
	return nil
}

// MarshalJSON marshals a JSON pointer.
func (p *JSONPointer) MarshalJSON() ([]byte, error) {
	return []byte(`"` + p.String() + `"`), nil
}
