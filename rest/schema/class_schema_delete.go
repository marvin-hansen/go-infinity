package schema

import (
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteClassSchema(classID string) *RequestForDeleteClassSchema {
	return &RequestForDeleteClassSchema{classID: classID, version: defaultVersion}
}

// RequestForDeleteClassSchema Deletes the class description for an existing type specified by fully qualified class name or class number.
// Before deleting a class, remove all objects of that class from the federated database. Objects of the class become inaccessible to applications after the class is deleted.
// Also, attributes of other classes that reference objects of the deleted class are deleted.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestSchemaTypeDelete.html%23
type RequestForDeleteClassSchema struct {
	classID string
	version string
}

func (req *RequestForDeleteClassSchema) Path() string {
	return "/" + req.version + "/" + "schema" + "/" + req.classID
}

func (req *RequestForDeleteClassSchema) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteClassSchema) Query() string {
	return ""
}

func (req *RequestForDeleteClassSchema) Payload() []byte {
	return nil
}

func (req *RequestForDeleteClassSchema) ResponseCode() int {
	return 204 //  No Content
}

//**// Response //**//

func NewResponseForDeleteClassSchema() *ResponseForDeleteClassSchema {
	return new(ResponseForDeleteClassSchema)
}

func (r ResponseForDeleteClassSchema) String() string {
	return ""
}

type ResponseForDeleteClassSchema struct{}
