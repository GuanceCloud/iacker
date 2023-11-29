package schema

import (
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/types"
	v1 "github.com/GuanceCloud/iacker/pkg/resource/v1"
	"github.com/go-kratos/kratos/v2/log"
)

// References is the references of the resource
type References struct {
	refs map[string][]*types.Identifier
}

// Add adds the reference
func (refs *References) Add(id *types.Identifier) {
	if refs.refs == nil {
		refs.refs = make(map[string][]*types.Identifier)
	}
	refs.refs[id.ResourceType] = append(refs.refs[id.ResourceType], id)
}

// Get gets the references
func (refs *References) Get(typeName string) []*types.Identifier {
	return refs.refs[typeName]
}

// Exists checks if the reference exists
func (refs *References) Exists(id *types.Identifier) bool {
	l, ok := refs.refs[id.ResourceType]
	if !ok {
		return false
	}
	for _, i := range l {
		if i.String() == id.String() {
			return true
		}
	}
	return false
}

// Types returns the types of the references
func (refs *References) Types() []string {
	typeNames := make([]string, 0, len(refs.refs))
	for t := range refs.refs {
		typeNames = append(typeNames, t)
	}
	return typeNames
}

type schemaRef struct {
	TypeName string
	Ref      string
}

func (s *rawSchema) genSchemaRefs(m *v1.Model) []*schemaRef {
	var refs []*schemaRef
	for _, prop := range m.Properties {
		if prop.Schema.Elem != nil && prop.Schema.Elem.GetRef() != "" {
			refs = append(refs, &schemaRef{
				TypeName: prop.Schema.Elem.GetRef(),
				Ref:      "/" + prop.Name,
			})
			continue
		}
		if prop.Schema.GetRef() != "" {
			refs = append(refs, &schemaRef{
				TypeName: prop.Schema.GetRef(),
				Ref:      "/" + prop.Name,
			})
			continue
		}

		if prop.Schema.GetType() == "object" {
			// get child model schema
			modelSchema, ok := s.schema.Models[prop.Schema.GetModel()]
			if !ok {
				log.Warnf("model not found: %s", prop.Schema.GetModel())
				continue
			}

			// generate references of the child model
			childRefs := s.genSchemaRefs(modelSchema)
			for _, ref := range childRefs {
				refs = append(refs, &schemaRef{
					TypeName: ref.TypeName,
					Ref:      "/" + prop.GetName() + ref.Ref,
				})
			}
			continue
		}
	}
	return refs
}
