package app

import (
	"database/sql"

	"github.com/stdyum/api-common/proto/impl/studyplaces"
	"github.com/stdyum/api-common/server"
	"github.com/stdyum/api-types-registry/internal/app/controllers"
	"github.com/stdyum/api-types-registry/internal/app/errors"
	"github.com/stdyum/api-types-registry/internal/app/handlers"
	"github.com/stdyum/api-types-registry/internal/app/repositories"
)

func New(database *sql.DB, studyPlacesClient studyplaces.StudyplacesClient) (server.Routes, controllers.Controller, error) {
	repo := repositories.New(database)

	ctrl := controllers.New(repo, studyPlacesClient)

	errors.Register()

	httpHndl := handlers.NewHTTP(ctrl)
	grpcHndl := handlers.NewGRPC(ctrl)

	routes := server.Routes{
		GRPC: grpcHndl,
		HTTP: httpHndl,
	}

	return routes, ctrl, nil
}
