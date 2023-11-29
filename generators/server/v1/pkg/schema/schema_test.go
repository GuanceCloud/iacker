package schema

import (
	"testing"

	types2 "github.com/GuanceCloud/iacker/generators/server/v1/pkg/types"
	"github.com/stretchr/testify/assert"
)

// Test Schema
func TestSchema(t *testing.T) {
	type testInput struct {
		TypeName string
		Schema   string
		Resource string
	}
	type testWant struct {
		Id   types2.Identifier
		Refs []types2.Identifier
	}
	tests := []struct {
		name    string
		input   testInput
		want    testWant
		wantErr bool
	}{
		{
			name: "OK",
			input: testInput{
				TypeName: "Test",
				Schema:   testSchemaJson,
				Resource: testResourceJson,
			},
			want: testWant{
				Id: types2.Identifier{
					ResourceType: "Test",
				},
				Refs: []types2.Identifier{
					{
						ResourceType: "Resource",
					},
					{
						ResourceType: "User",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewFromBytes(tt.input.TypeName, []byte(tt.input.Schema))
			if !assert.NoError(t, err) {
				return
			}

			r := &types2.Resource{
				State: tt.input.Resource,
			}

			// check the references
			refs, err := s.GenRefs(r)
			if !assert.NoError(t, err) {
				return
			}
			for _, refId := range tt.want.Refs {
				assert.True(t, refs.Exists(&refId))
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
	"ref": "resource-id",
	"user": {
		"id": "user-id"
	}
}`
