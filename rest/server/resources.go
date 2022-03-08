package server

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForResources(version string) *RequestForResources {
	return &RequestForResources{version: version}
}

// RequestForResources Retrieves information about the available resources for this version of the API.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestVersionGET.html%23
type RequestForResources struct {
	version string
}

func (req *RequestForResources) Path() string {
	return "/" + req.version
}

func (req *RequestForResources) Method() string {
	return http.MethodGet
}

func (req *RequestForResources) Query() string {
	return ""
}

func (req *RequestForResources) Payload() []byte {
	return nil
}

func (req *RequestForResources) ResponseCode() int {
	return 200
}

//**// Response //**//

func NewResponseForResources() *ResponseForResources {
	return &ResponseForResources{}
}

func (r ResponseForResources) String() string {
	return fmt.Sprintf("[Version: %v, Resources: %v]", r.Version, r.Resources)
}

func (r *ResponseForResources) GetRawMessage() []byte {
	return r.RawMessage
}

func (r *ResponseForResources) SetRawMessage(raw []byte) {
	r.RawMessage = raw
}

type ResponseForResources RestResources

type RestResources struct {
	Version    string      `json:"version"`
	Resources  []*Resource `json:"resources"`
	RawMessage []byte
}

type Resource struct {
	Uri              string   `json:"uri"`
	SupportedMethods []string `json:"supportedMethods"`
}

func (r Resource) String() string {
	return fmt.Sprintf("[URI: %v, Supported Methods: %v]", r.Uri, r.SupportedMethods)
}
