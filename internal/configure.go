package internal

import (
	"github.com/stdyum/api-common/grpc/clients"
	"github.com/stdyum/api-common/server"
	"github.com/stdyum/api-types-registry/internal/app"
	"github.com/stdyum/api-types-registry/internal/app/controllers"
	"github.com/stdyum/api-types-registry/internal/config"
)

func Configure() (server.Routes, controllers.Controller, error) {
	db, err := config.ConnectToDatabase(config.Config.Database)
	if err != nil {
		return server.Routes{}, nil, err
	}

	studyPlacesServer, err := config.ConnectToStudyPlacesServer(config.Config.StudyPlacesGRpc)
	if err != nil {
		return server.Routes{}, nil, err
	}
	clients.StudyPlacesGRpcClient = studyPlacesServer

	routes, ctrl, err := app.New(db, studyPlacesServer)
	if err != nil {
		return server.Routes{}, nil, err
	}

	routes.Ports = config.Config.Ports

	return routes, ctrl, nil
}
