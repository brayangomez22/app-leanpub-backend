//go:build wireinject
// +build wireinject

package test

import "github.com/google/wire"

func CreateApp() *Application {
	wire.Build(DbGateweyProvider, TestApplicacion)
	return new(Application)
}
