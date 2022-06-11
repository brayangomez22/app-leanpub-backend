package app

import (
	"github.com/gorilla/mux"
	"leanpub-app/domain"
	"leanpub-app/domain/usecases"
)

type Application struct {
	Router       *mux.Router
	datastore    domain.DatabaseGateway
	userUseCases usecases.UserUseCase
}

func NewApplication(datastore domain.DatabaseGateway, userUseCase usecases.UserUseCase) *Application {
	return &Application{
		datastore:    datastore,
		userUseCases: userUseCase,
	}
}
