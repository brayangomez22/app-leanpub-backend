// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"leanpub-app/domain/usecases"
	"leanpub-app/infra/datastore"
)

// Injectors from wire.go:

func CreateApp() *Application {
	databaseGateway := datastore.NewMongoGatewayImpl()
	userUseCase := usecases.NewUserUseCase(databaseGateway)
	application := NewApplication(databaseGateway, userUseCase)
	return application
}
