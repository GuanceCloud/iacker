package types

import (
	"encoding/json"
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"
)

// ResourcePatch is the actual representation of a resource patch.
type ResourcePatch struct {
	// Id is the identifier of the resource to be patched.
	Id *Identifier

	// Patches is a list of JSON patches.
	Patches json.RawMessage
}

// Apply applies the patch to the resource.
func (g *ResourcePatch) Apply(rs *Resource) (*Resource, error) {
	patch, err := jsonpatch.DecodePatch(g.Patches)
	if err != nil {
		return nil, fmt.Errorf("failed to decode patch: %w", err)
	}

	state := stateWrapper{Attributes: json.RawMessage(rs.State)}
	rsBytes, err := json.Marshal(state)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource: %w", err)
	}

	result, err := patch.ApplyIndent(rsBytes, "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to apply patch: %w", err)
	}
	rs.State = string(result)
	return rs, nil
}

type stateWrapper struct {
	Attributes json.RawMessage `json:"attributes"`
}
