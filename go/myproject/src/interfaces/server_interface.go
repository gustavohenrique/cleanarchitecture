package interfaces

type IServer interface {
	New(services IService) IServer
	GetRawServer() interface{}
	Configure(params interface{})
	Start(address string, port int) error
}
