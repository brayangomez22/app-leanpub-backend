package app

import (
	"github.com/google/wire"
	"leanpub-app/domain/usecases"
	"leanpub-app/infra/datastore"
)

var DataStoreProvider = wire.NewSet(datastore.NewMongoGatewayImpl)
var UserUseCasesProvider = wire.NewSet(usecases.NewUserUseCase)
var BookUseCasesProvider = wire.NewSet(usecases.NewBookUseCase)
var ShoppingCartUseCasesProvider = wire.NewSet(usecases.NewShoppingCartUseCase)
var AppProvider = wire.NewSet(NewApplication)
