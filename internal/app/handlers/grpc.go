package handlers

import (
	"github.com/stdyum/api-common/grpc"
	"github.com/stdyum/api-common/proto/impl/types_registry"
	"github.com/stdyum/api-types-registry/internal/app/controllers"
)

type GRPC interface {
	grpc.Routes
	types_registry.TypesRegistryServer
}

type gRPC struct {
	types_registry.UnsafeTypesRegistryServer

	controller controllers.Controller
}

func NewGRPC(controller controllers.Controller) GRPC {
	return &gRPC{
		controller: controller,
	}
}
