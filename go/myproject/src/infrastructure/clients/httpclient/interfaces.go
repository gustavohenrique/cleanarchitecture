package httpclient

type HttpClient interface {
	Failed(err error, statusCode int) bool
	PATCH(url string, data []byte) ([]byte, int, error)
	DELETE(url string) ([]byte, int, error)
	PUT(url string, data []byte) ([]byte, int, error)
	POST(url string, data []byte) ([]byte, int, error)
	GET(url string) ([]byte, int, error)
}
