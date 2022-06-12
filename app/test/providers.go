package test

import (
	"github.com/google/wire"
	"leanpub-app/domain"
)

var DbGateweyProvider = wire.NewSet(NewDbGateway, wire.Bind(new(domain.DatabaseGateway), new(DbGateway)))
var TestApplicacion = wire.NewSet(NewApplication)