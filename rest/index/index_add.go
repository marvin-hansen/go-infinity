package index

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForAddIndex(jsonIndex string) *RequestForAddIndex {
	return &RequestForAddIndex{payload: []byte(jsonIndex), version: defaultVersion}
}

// RequestForAddIndex Creates an index for the given class using the given attributes as keys.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestIndexPOST.html%23
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
	Name       string `json:"name"`
	Uri        string `json:"uri"`
	RawMessage []byte
}

func (r *ResponseForAddIndex) GetRawMessage() []byte {
	return r.RawMessage
}

func (r *ResponseForAddIndex) SetRawMessage(raw []byte) {
	r.RawMessage = raw
}
