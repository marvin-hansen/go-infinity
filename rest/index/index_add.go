package index

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForAddIndex(index []byte) *RequestForAddIndex {
	return &RequestForAddIndex{payload: index, version: defaultVersion}
}

type RequestForAddIndex struct {
	payload []byte
	version string
}

func (req *RequestForAddIndex) Path() string {
	return "/" + req.version + "/index"
}

func (req *RequestForAddIndex) Method() string {
	return http.MethodPost
}

func (req *RequestForAddIndex) Query() string {
	return ""
}

func (req *RequestForAddIndex) Payload() []byte {
	return req.payload
}

func (req *RequestForAddIndex) ResponseCode() int {
	return 201 // created
}

//**// Response //**//

func NewResponseForAddIndex() *ResponseForAddIndex {
	return new(ResponseForAddIndex)
}

func (r ResponseForAddIndex) String() string {
	return fmt.Sprintf("Name: %v \n, Uri: %v \n", r.Name, r.Uri)
}

type ResponseForAddIndex struct {
	Name       string
	Uri        string
	RawMessage []byte
}

func (r *ResponseForAddIndex) GetRawMessage() []byte {
	return r.RawMessage
}

func (r *ResponseForAddIndex) SetRawMessage(raw []byte) {
	r.RawMessage = raw
}
