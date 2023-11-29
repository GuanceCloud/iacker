package schema

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/go-multierror"

	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/types"

	v1 "github.com/GuanceCloud/iacker/pkg/resource/v1"
)

// Schema is the interface for resource schema
type Schema interface {
	Spec() *v1.Resource
	GenRefs(rs *types.Resource) (*References, error)
	Validate(rs *types.Resource) error
	Decode(state string, out interface{}) error
	Encode(in interface{}) (string, error)
}

// Compile creates a new schema from raw json
func Compile(raw string, typeName string) (Schema, error) {
	rs := &v1.Resource{}
	err := json.Unmarshal([]byte(raw), rs)
	if err != nil {
		return nil, err
	}
	return &rawSchema{
		typeName: typeName,
		raw:      raw,
		schema:   rs,
	}, nil
}

// MustCompile creates a new schema from raw json
func MustCompile(raw string, typeName string) Schema {
	s, err := Compile(raw, typeName)
	if err != nil {
		log.Fatal(fmt.Errorf("compile schema %q failed: %w", typeName, err))
	}
	return s
}

// rawSchema is the base schema
type rawSchema struct {
	typeName string
	raw      string
	schema   *v1.Resource
}

// Spec returns the resource schema
func (s *rawSchema) Spec() *v1.Resource {
	return s.schema
}

// ResourceType returns the type name of Workspace
func (s *rawSchema) ResourceType() string {
	return s.typeName
}

// GenRefs generates the references of the resource
func (s *rawSchema) GenRefs(rs *types.Resource) (*References, error) {
	var mErr error
	refs := &References{refs: make(map[string][]*types.Identifier)}
	for _, ref := range s.genSchemaRefs(s.schema.Models[s.typeName]) {
		paths := strings.Split(strings.TrimPrefix(ref.Ref, "/"), "/")
		pointer := NewJSONPointer(paths...)
		id, err := types.ParseIdentifier(pointer.Get(rs.State))
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("parse id failed: %w", err))
		}
		refs.Add(&id)
	}
	return refs, nil
}

// Validate validates the resource
// TODO: implement validate by cue-lang
func (s *rawSchema) Validate(rs *types.Resource) error {
	return nil
}

// Decode decodes the resource state to the given struct
func (s *rawSchema) Decode(state string, v interface{}) error {
	return json.Unmarshal([]byte(state), v)
}

// Encode encodes the given struct to the resource state
func (s *rawSchema) Encode(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}
