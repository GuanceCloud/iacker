package schema

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test JSONPointer
func TestJSONPointer(t *testing.T) {
	tests := []struct {
		name   string
		ref    string
		state  string
		expect string
	}{
		{
			name:   "OK",
			ref:    "/user/id",
			state:  `{"user": {"id": 123}}`,
			expect: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var pointer *JSONPointer
			if err := json.Unmarshal([]byte(fmt.Sprintf(`"%s"`, tt.ref)), &pointer); err != nil {
				assert.NoError(t, err)
				return
			}
			if marshaled, err := json.Marshal(pointer); err != nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, fmt.Sprintf(`"%s"`, tt.ref), string(marshaled))
			}
			assert.Equal(t, tt.ref, pointer.String())
			assert.Equal(t, tt.expect, pointer.Get(tt.state))
		})
	}
}
