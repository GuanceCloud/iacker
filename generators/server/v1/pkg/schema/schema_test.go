package schema

import (
	"testing"

	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/types"
	"github.com/stretchr/testify/assert"
)

// Test Schema
func TestSchema(t *testing.T) {
	type testInput struct {
		TypeName string
		Schema   string
		Resource string
	}
	tests := []struct {
		name    string
		input   testInput
		want    []string
		wantErr bool
	}{
		{
			name: "OK",
			input: testInput{
				TypeName: "Test",
				Schema:   testSchemaJson,
				Resource: testResourceJson,
			},
			want: []string{
				"grn:1:2:3:4:5:6",
				"grn:1:2:3:4:5:42",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := Compile(tt.input.Schema, tt.input.TypeName)
			if !assert.NoError(t, err) {
				return
			}
			r := &types.Resource{
				State: tt.input.Resource,
			}

			// check the references
			refs, err := s.GenRefs(r)
			if !assert.NoError(t, err) {
				return
			}
			for _, refId := range tt.want {
				id, err := types.ParseIdentifier(refId)
				if assert.NoError(t, err) {
					assert.True(t, refs.Exists(&id))
				}
			}
		})
	}
}

const testSchemaJson = `{
	"plural": "Test",
	"models": {
		"Test": {
			"properties": [
				{
					"name": "id",
					"schema": {
						"type": "string"
					}
				},
				{
					"name": "ref",
					"schema": {
						"type": "ref",
						"ref": "Resource"
					}
				},
				{
					"name": "user",
					"schema": {
						"type": "object",
						"model": "User"
					}
				}
            ]
		},
		"User": {
			"properties": [
				{
					"name": "id",
					"schema": {
						"type": "ref",
						"ref": "User"
					}
				}
			]
		}
	},
	"identifier": {
		"primary": [
			"/properties/id"
		]
	}
}`

const testResourceJson = `{
	"id": "test-id",
	"ref": "grn:1:2:3:4:5:6",
	"user": {
		"id": "grn:1:2:3:4:5:42"
	}
}`
