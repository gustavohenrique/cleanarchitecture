package servers

type Server interface {
	RawServer() interface{}
	Configure(params ...interface{})
	Start() error
}
