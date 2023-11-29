package server

import (
	v12 "github.com/GuanceCloud/iacker/generators/server/v1/pkg/config/v1"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "github.com/GuanceCloud/iacker/generators/server/v1/api/cloudcontrol/v1"
	"github.com/GuanceCloud/iacker/generators/server/v1/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *v12.Server, svc *service.CloudControlService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != time.Duration(0) {
		opts = append(opts, http.Timeout(c.Http.Timeout))
	}
	srv := http.NewServer(opts...)
	v1.RegisterCloudControlServiceHTTPServer(srv, svc)
	return srv
}
