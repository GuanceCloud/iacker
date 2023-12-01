package server

import (
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"github.com/GuanceCloud/iacker/generators/server/v1/internal/biz"
	"github.com/GuanceCloud/iacker/generators/server/v1/internal/service"
	ib "github.com/GuanceCloud/iacker/generators/server/v1/pkg/backend"
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/config/v1"
)

type Server interface {
	Run() error
}

func NewServer(b ib.Backend, cfg *config.Server) Server {
	instInfo := cfg.Instance

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", instInfo.ID,
		"service.name", instInfo.Name,
		"service.version", instInfo.Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	cloudControlUseCase := biz.NewCloudControlUseCase(b, logger)
	cloudControlService := service.NewCloudControlService(cloudControlUseCase)
	grpcServer := NewGRPCServer(cfg, cloudControlService, logger)
	httpServer := NewHTTPServer(cfg, cloudControlService, logger)

	return kratos.New(
		kratos.ID(instInfo.ID),
		kratos.Name(instInfo.Name),
		kratos.Version(instInfo.Version),
		kratos.Metadata(instInfo.Metadata),
		kratos.Logger(logger),
		kratos.Server(
			grpcServer,
			httpServer,
		),
	)
}
