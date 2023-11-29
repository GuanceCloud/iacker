package server

import (
	"github.com/GuanceCloud/iacker/generators/server/v1/internal/server"
	"github.com/GuanceCloud/iacker/generators/server/v1/pkg/config/v1"
)

type RunOptions = server.RunOptions

type Config = config.Server

func Run(opts RunOptions) error {
	return server.Run(opts)
}
