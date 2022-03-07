package schema

import (
	"net/http"
)

//**// Request //**//

func NewRequestForGetSchema() *RequestForGetSchema {
	return &RequestForGetSchema{defaultVersion}
}

// RequestForGetSchema Retrieves a representation of the complete schema for a federated database.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestSchemaGET.html%23
type RequestForGetSchema struct {
	version string
}

func (req *RequestForGetSchema) Path() string {
	return "/" + req.version + "/schema"
}

func (req *RequestForGetSchema) Method() string {
	return http.MethodGet
}

func (req *RequestForGetSchema) Query() string {
	return ""
}

func (req *RequestForGetSchema) Payload() []byte {
	return nil
}

func (req *RequestForGetSchema) ResponseCode() int {
	return 200
}

//**// Response //**//

func NewResponseForGetSchema() *ResponseForRetrieveSchema {
	return new(ResponseForRetrieveSchema)
}

type ResponseForRetrieveSchema []Schema
