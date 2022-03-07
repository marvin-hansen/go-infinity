package server

import "net/http"

func NewRequestForShutdown() *RequestForShutdown {
	return new(RequestForShutdown)
}

// RequestForShutdown Terminates the InfiniteGraph REST server running at the provided URL.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestVersionShutdownPOST.html%23
type RequestForShutdown struct{}

func (req *RequestForShutdown) Path() string {
	return "/shutdown"
}

func (req *RequestForShutdown) Method() string {
	return http.MethodPost
}

func (req *RequestForShutdown) Query() string {
	return ""
}

func (req *RequestForShutdown) Payload() []byte {
	return nil
}

func (req *RequestForShutdown) ResponseCode() int {
	return 204 // No Content
}

//**// Response //**//

func NewResponseForShutdown() *ResponseForShutdown {
	return new(ResponseForShutdown)
}

// ResponseForShutdown actually returns 204 No Content from the REST API thus empty struct.
type ResponseForShutdown struct{}
