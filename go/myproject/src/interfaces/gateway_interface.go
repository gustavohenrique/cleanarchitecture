package interfaces

type IGateway interface {
	Inject(ds IDataStore) IGateway
	TodoGateway() ITodoGateway
}
