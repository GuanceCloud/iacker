package testing

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/provider"
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/types"
)

type ResourceTestOptions struct {
	Context     *provider.Context
	Middlewares []provider.Middleware
}

type ResourceTestCase struct {
	// Name is the name of the test case.
	Name string

	// ConfigureFunc returns the setup options.
	ConfigureFunc func() *ResourceTestOptions

	// Resource is the resource to test.
	Resource provider.Resource

	// State is the desired state
	State string
}

func (tc ResourceTestCase) prepare(t *testing.T) context.Context {
	var err error
	opts := tc.ConfigureFunc()
	iackerContext := opts.Context
	iackerContext.TypeName = tc.Resource.GetConfig().TypeName
	ctx := provider.WithContext(context.Background(), iackerContext)
	for _, m := range opts.Middlewares {
		ctx, err = m.Prepare(ctx)
		if err != nil {
			t.Fatalf("prepare middleware: %v", err)
		}
	}
	return ctx
}

type testStep struct {
	name string
	fn   func(t *testing.T) error
}

func RunTestResource(t *testing.T, tc *ResourceTestCase) {
	ctx := tc.prepare(t)

	state := &provider.State{
		Config: tc.Resource.GetConfig(),
		State:  tc.State,
	}

	steps := []testStep{
		{
			name: "validate",
			fn: func(t *testing.T) error {
				return tc.Resource.Validate(ctx, state)
			},
		},
		{
			name: "create",
			fn: func(t *testing.T) error {
				return tc.Resource.Create(ctx, state)
			},
		},
		{
			name: "delete",
			fn: func(t *testing.T) error {
				return tc.Resource.Delete(ctx, state)
			},
		},
	}

	for _, step := range steps {
		ok := t.Run(fmt.Sprintf("%s/%s", tc.Name, step.name), func(t *testing.T) {
			if err := step.fn(t); err != nil {
				t.Fatalf("%s: %v", step.name, err)
			}
		})
		if !ok {
			break
		}
	}
}

func RunTestDataSource(t *testing.T, tc *ResourceTestCase) {
	ctx := tc.prepare(t)

	t.Run(fmt.Sprintf("%s/list", tc.Name), func(t *testing.T) {
		result, err := tc.Resource.List(ctx, &types.ResourceQuery{
			TypeName:   tc.Resource.GetConfig().TypeName,
			MaxResults: 100,
		})
		if err != nil {
			t.Fatalf("read: %v", err)
		}
		assert.NotEmpty(t, result.Items)
	})
}
