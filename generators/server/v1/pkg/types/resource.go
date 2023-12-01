package types

import "time"

// Result is the result of a resource query.
type Result struct {
	Items     []*Resource
	NextToken string
}

// Resource is the actual representation of a resource.
type Resource struct {
	Identifier *Identifier
	State      string
	CreatedAt  time.Time
}
