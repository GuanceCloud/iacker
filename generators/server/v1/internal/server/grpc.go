package server

import (
	v12 "github.com/GuanceCloud/iacker/generators/server/v1/pkg/config/v1"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	v1 "github.com/GuanceCloud/iacker/generators/server/v1/api/cloudcontrol/v1"
	"github.com/GuanceCloud/iacker/generators/server/v1/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *v12.Server, svc *service.CloudControlService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != time.Duration(0) {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterCloudControlServiceServer(srv, svc)
	return srv
}
