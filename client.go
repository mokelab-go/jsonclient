package jsonclient

const (
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PUT    = "PUT"
	METHOD_DELETE = "DELETE"
)

type Client interface {
	SetURL(url string)
	SetMethod(method string)
	SetHeader(key, value string)
	Send(body map[string]interface{}) (Response, error)
}

type Response struct {
	Status  int
	Headers map[string]string
	Body    map[string]interface{}
}
