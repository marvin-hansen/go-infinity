package schema

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForUpdateSchema(schema []byte) *RequestForUpdateSchema {
	return &RequestForUpdateSchema{schema, defaultVersion}
}

// RequestForUpdateSchema Provides one or more class descriptions, each of which results in the creation of a new class or the addition of new attributes to an existing class.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestSchemaPUT.html%23
type RequestForUpdateSchema struct {
	payload []byte
	version string
}

func (req *RequestForUpdateSchema) Path() string {
	return "/" + req.version + "/schema"
}

func (req *RequestForUpdateSchema) Method() string {
	return http.MethodPut
}

func (req *RequestForUpdateSchema) Query() string {
	return ""
}

func (req *RequestForUpdateSchema) Payload() []byte {
	return req.payload
}

//**// Response //**//

func NewResponseForUpdateSchema() *ResponseForUpdateSchema {
	return new(ResponseForUpdateSchema)
}

func (r ResponseForUpdateSchema) String() string {
	return fmt.Sprintf("Class Numbers: %v \n Uris: %v", r.ClassNumbers, r.Uris)
}

type ResponseForUpdateSchema UpdateSchema

type UpdateSchema struct {
	ClassNumbers []string `json:"classNumbers"`
	Uris         []string `json:"uris"`
}
