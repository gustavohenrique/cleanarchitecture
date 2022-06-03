package server

type Server interface {
	GetRawServer() interface{}
	Configure(params interface{})
	Start(address string, port int) error
}
