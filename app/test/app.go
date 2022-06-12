package test

type Application struct {
	DataStore DbGateway
}

func NewApplication(datastoreGateway DbGateway)	*Application {
	return &Application{
		DataStore: datastoreGateway,
	}
}