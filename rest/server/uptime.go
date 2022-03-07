package server

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForUptime() *RequestForUptime {
	return new(RequestForUptime)
}

// RequestForUptime Retrieves information about the InfiniteGraph REST server, including how long it has been running and the available versions of the API.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestGET.html%23
type RequestForUptime struct{}

func (req *RequestForUptime) Path() string {
	return "/"
}

func (req *RequestForUptime) Method() string {
	return http.MethodGet
}

func (req *RequestForUptime) Query() string {
	return ""
}

func (req *RequestForUptime) Payload() []byte {
	return nil
}

func (req *RequestForUptime) ResponseCode() int {
	return 200
}

//**// Response //**//

func NewResponseForUptime() *ResponseForUptime {
	return new(ResponseForUptime)
}

func (r ResponseForUptime) String() string {
	return fmt.Sprintf("[%v, %v]", r.Uptime, r.Versions)
}

type ResponseForUptime Uptime

type Uptime struct {
	Uptime   int64    `json:"uptime"`
	Versions []string `json:"versions"`
}
