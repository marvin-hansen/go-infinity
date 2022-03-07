package object

import (
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteObject(objectID string) *RequestForDeleteObject {
	return &RequestForDeleteObject{objectID: objectID, version: defaultVersion}
}

// RequestForDeleteObject Deletes the object with the given OID from the federated database.
// Does not return anything.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestObjectOidDELETE.html%23
type RequestForDeleteObject struct {
	objectID string
	version  string
}

func (req *RequestForDeleteObject) Path() string {
	return "/" + req.version + "/" + "object" + "/" + req.objectID
}

func (req *RequestForDeleteObject) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteObject) Query() string {
	return ""
}

func (req *RequestForDeleteObject) Payload() []byte {
	return nil
}

func (req *RequestForDeleteObject) ResponseCode() int {
	return 204 // No content
}

//**// Response //**//

// No response because of 204 status code
