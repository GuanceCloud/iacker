package config

import "time"

type Config struct {
	Backend *Backend
	Server  *Server
}

type Backend struct {
	Driver        string
	Source        string
	AutoMigration bool
}

type Server struct {
	Instance *Instance
	Http     *HttpServer
	Grpc     *GrpcServer
}

type HttpServer struct {
	Network string
	Addr    string
	Timeout time.Duration
}

type GrpcServer struct {
	Network string
	Addr    string
	Timeout time.Duration
}

type Instance struct {
	ID       string
	Name     string
	Version  string
	Metadata map[string]string
}
