package object

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForAddObject(jsonObject []byte) *RequestForAddObject {
	return &RequestForAddObject{payload: jsonObject, version: defaultVersion}
}

// RequestForAddObject Adds a new object to a federated database and assigns an OID.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestObjectPOST.html%23
type RequestForAddObject struct {
	payload []byte
	version string
}

func (req *RequestForAddObject) Path() string {
	return "/" + req.version + "/" + "object"
}

func (req *RequestForAddObject) Method() string {
	return http.MethodPost
}

func (req *RequestForAddObject) Query() string {
	return ""
}

func (req *RequestForAddObject) Payload() []byte {
	return req.payload
}

func (req *RequestForAddObject) ResponseCode() int {
	return 201 // created
}

//**// Response //**//

func NewResponseForAddObject() *ResponseForAddObject {
	return new(ResponseForAddObject)
}

func (r ResponseForAddObject) String() string {
	return fmt.Sprintf("ObjectID: %v, \n Uri: %v \n",
		r.ObjectID,
		r.Uri,
	)
}

type ResponseForAddObject struct {
	ObjectID   string `json:"__identifier__"`
	Uri        string `json:"uri"`
	RawMessage []byte
}

func (r *ResponseForAddObject) GetRawMessage() []byte {
	return r.RawMessage
}

func (r *ResponseForAddObject) SetRawMessage(raw []byte) {
	r.RawMessage = raw
}
