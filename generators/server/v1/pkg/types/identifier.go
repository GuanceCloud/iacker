package types

import (
	"fmt"
	"strings"
)

type Identifier struct {
	Partition    string
	Service      string
	Region       string
	Owner        string
	ResourceType string
	ResourceId   string
}

func (id Identifier) String() string {
	return fmt.Sprintf("grn:%s:%s:%s:%s:%s:%s", id.Partition, id.Service, id.Region, id.Owner, id.ResourceType, id.ResourceId)
}

// ParseIdentifier parses an identifier from a string.
func ParseIdentifier(id string) (Identifier, error) {
	tokens := strings.Split(id, ":")
	if len(tokens) != 7 {
		return Identifier{}, fmt.Errorf("invalid identifier: %s", id)
	}
	return Identifier{
		Partition:    tokens[1],
		Service:      tokens[2],
		Region:       tokens[3],
		Owner:        tokens[4],
		ResourceType: tokens[5],
		ResourceId:   tokens[6],
	}, nil
}
