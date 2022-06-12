//go:build wireinject
// +build wireinject

package app

import "github.com/google/wire"

func CreateApp() *Application {

	wire.Build(
		DataStoreProvider,
		UserUseCasesProvider,
		BookUseCasesProvider,
		AppProvider,
	)

	return new(Application)
}
