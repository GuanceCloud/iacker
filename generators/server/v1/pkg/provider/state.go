package provider

import (
	"fmt"

	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/types"
)

// State is the state of resource.
type State struct {
	// Config is the options of resource reconciler.
	Config *ResourceConfig

	// Identifier is the id of resource.
	// It is the id of resource in the resource management server.
	Identifier *types.Identifier

	// State is the state of resource.
	State string
}

// Decode decodes the state.
func (s *State) Decode(out interface{}) error {
	return s.Config.Schema.Decode(s.State, out)
}

// Encode encodes the state.
func (s *State) Encode(in interface{}) error {
	state, err := s.Config.Schema.Encode(in)
	if err != nil {
		return fmt.Errorf("encode state: %w", err)
	}
	s.State = state
	return nil
}
