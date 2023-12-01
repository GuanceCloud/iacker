package server

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hashicorp/go-multierror"

	"github.com/GuanceCloud/iacker/generators/server/v1/internal/backend"
	ib "github.com/GuanceCloud/iacker/generators/server/v1/pkg/backend"
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/config/v1"
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/provider"
)

type RunOptions struct {
	// Server config
	Config *config.Server

	// Backend config
	ResourceRepo ib.ResourceRepo
	RequestRepo  ib.RequestRepo
	Middlewares  []provider.Middleware

	// Resources config
	Resources map[string]provider.Factory
}

func Run(opts RunOptions) error {
	var mErr error
	resources := map[string]provider.Resource{}

	for name, factor := range opts.Resources {
		r, err := factor(log.DefaultLogger)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("create resource %s failed: %w", name, err))
			continue
		}
		resources[name] = r
	}
	if mErr != nil {
		return mErr
	}

	b := &backend.StateBackend{
		Request:     opts.RequestRepo,
		Resource:    opts.ResourceRepo,
		Resources:   resources,
		Middlewares: opts.Middlewares,
	}
	return NewServer(b, opts.Config).Run()
}
