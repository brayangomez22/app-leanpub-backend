
package app

import (
	"github.com/google/wire"
	"leanpub-app/domain/usecases"
	"leanpub-app/infra/datastore"
)

var DataStoreProvider = wire.NewSet(datastore.NewMongoGatewayImpl)
var UserUseCasesProvider = wire.NewSet(usecases.NewUserUseCase)
var AppProvider = wire.NewSet(NewApplication)
