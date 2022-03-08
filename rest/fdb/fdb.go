package fdb

import (
	"fmt"
	"go-infinity/rest/flags"
	"net/http"
)

const defaultVersion = flags.VERSION

//**// Request //**//

func NewRequestForBootPath() *RequestForBootPath {
	return &RequestForBootPath{defaultVersion}
}

func NewRequestForBootPathApiVersion(apiVersion string) *RequestForBootPath {
	return &RequestForBootPath{apiVersion}
}

// RequestForBootPath Retrieves the path to the boot file for the current federated database.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestFdGET.html%23wwconnect_header
type RequestForBootPath struct {
	version string
}

func (req *RequestForBootPath) Path() string {
	return "/" + req.version + "/fd"
}

func (req *RequestForBootPath) Method() string {
	return http.MethodGet
}

func (req *RequestForBootPath) Query() string {
	return ""
}

func (req *RequestForBootPath) Payload() []byte {
	return nil
}

func (req *RequestForBootPath) ResponseCode() int {
	return 200
}

//**// Response //**//

func NewResponseForBootPath() *ResponseForBootPath {
	return new(ResponseForBootPath)
}

func (r *ResponseForBootPath) String() string {
	return fmt.Sprintf("Bootfile: %v", r.Bootfile)
}

func (r *ResponseForBootPath) GetRawMessage() []byte {
	return r.RawMessage
}

func (r *ResponseForBootPath) SetRawMessage(raw []byte) {
	r.RawMessage = raw
}

type ResponseForBootPath Bootfile

type Bootfile struct {
	Bootfile   string `json:"bootfile"`
	RawMessage []byte
}
