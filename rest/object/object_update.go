package object

import (
	"net/http"
)

//**// Request //**//

func NewRequestForUpdateObject(objectID string, jsonObject []byte) *RequestForUpdateObject {
	return &RequestForUpdateObject{objectID: objectID, payload: jsonObject, version: defaultVersion}
}

// RequestForUpdateObject Modifies an existing object by changing attribute values.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestObjectOidPUT.html%23
type RequestForUpdateObject struct {
	objectID string
	payload  []byte
	version  string
}

func (req *RequestForUpdateObject) Path() string {
	return "/" + req.version + "/" + "object" + "/" + req.objectID

}

func (req *RequestForUpdateObject) Method() string {
	return http.MethodPut
}

func (req *RequestForUpdateObject) Query() string {
	return ""
}

func (req *RequestForUpdateObject) Payload() []byte {
	return req.payload
}

func (req *RequestForUpdateObject) ResponseCode() int {
	return 204 //No content
}
