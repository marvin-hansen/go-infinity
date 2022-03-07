package schema

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForAddSchema(schema []byte) *RequestForAddSchema {
	return &RequestForAddSchema{schema, defaultVersion}
}

// RequestForAddSchema Adds one or more new class descriptions to the schema of a federated database.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestSchemaPOST.html%23
type RequestForAddSchema struct {
	payload []byte
	version string
}

func (req *RequestForAddSchema) Path() string {
	return "/" + req.version + "/schema"
}

func (req *RequestForAddSchema) Method() string {
	return http.MethodPost
}

func (req *RequestForAddSchema) Query() string {
	return ""
}

func (req *RequestForAddSchema) Payload() []byte {
	return req.payload
}

//**// Response //**//

func NewResponseForAddSchema() *ResponseForAddSchema {
	return new(ResponseForAddSchema)
}

func (r ResponseForAddSchema) String() string {
	return fmt.Sprintf("Class Numbers: %v \n Uris: %v", r.ClassNumbers, r.Uris)
}

type ResponseForAddSchema AddSchema

type AddSchema struct {
	ClassNumbers []string `json:"classNumbers"`
	Uris         []string `json:"uris"`
}
