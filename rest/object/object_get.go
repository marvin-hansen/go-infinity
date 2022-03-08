package object

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetObject(objectID string) *RequestForGetObject {
	return &RequestForGetObject{objectID: objectID, version: defaultVersion}
}

// RequestForGetObject Retrieves a representation of the object with the given OID.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestObjectOidGET.html%23
type RequestForGetObject struct {
	objectID string
	version  string
}

func (req *RequestForGetObject) Path() string {
	return "/" + req.version + "/" + "object" + "/" + req.objectID
}

func (req *RequestForGetObject) Method() string {
	return http.MethodGet
}

func (req *RequestForGetObject) Query() string {
	return ""
}

func (req *RequestForGetObject) Payload() []byte {
	return nil
}

func (req *RequestForGetObject) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetObject() *ResponseForGetObject {
	return new(ResponseForGetObject)
}

func (r ResponseForGetObject) String() string {
	return fmt.Sprintf("Class: %v \n  RawMessage: %v \n  ",
		r.Class,
		string(r.RawMessage),
	)
}

type ResponseForGetObject struct {
	Class      string `json:"__class__"`
	RawMessage []byte
}

func (r *ResponseForGetObject) GetRawMessage() []byte {
	return r.RawMessage
}

func (r *ResponseForGetObject) SetRawMessage(raw []byte) {
	r.RawMessage = raw
}
