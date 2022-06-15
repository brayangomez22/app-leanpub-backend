package app

import (
	"github.com/gorilla/mux"
	"leanpub-app/domain"
	"leanpub-app/domain/usecases"
)

type Application struct {
	Router               *mux.Router
	datastore            domain.DatabaseGateway
	userUseCases         usecases.UserUseCase
	bookUseCases         usecases.BookUseCase
	shoppingCartUseCases usecases.ShoppingCartUseCase
}

func NewApplication(
	datastore domain.DatabaseGateway,
	userUseCase usecases.UserUseCase,
	bookUseCases usecases.BookUseCase,
	shoppingCartUseCases usecases.ShoppingCartUseCase,
) *Application {
	return &Application{
		datastore:            datastore,
		userUseCases:         userUseCase,
		bookUseCases:         bookUseCases,
		shoppingCartUseCases: shoppingCartUseCases,
	}
}
