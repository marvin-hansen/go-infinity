package schema

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForUpdateClassSchema(classID string, classSchema []byte) *RequestForUpdateClassSchema {
	return &RequestForUpdateClassSchema{classID: classID, payload: classSchema, version: defaultVersion}
}

type RequestForUpdateClassSchema struct {
	classID string
	payload []byte
	version string
}

func (req *RequestForUpdateClassSchema) Path() string {
	return "/" + req.version + "/" + "schema" + "/" + req.classID
}

func (req *RequestForUpdateClassSchema) Method() string {
	return http.MethodPut
}

func (req *RequestForUpdateClassSchema) Query() string {
	return ""
}

func (req *RequestForUpdateClassSchema) Payload() []byte {
	return nil
}

func (req *RequestForUpdateClassSchema) ResponseCode() int {
	return 201 // Created
}

//**// Response //**//

func NewResponseForUpdateClassSchema() *ResponseForUpdateClassSchema {
	return new(ResponseForUpdateClassSchema)
}

func (r ResponseForUpdateClassSchema) String() string {
	return fmt.Sprintf("Class Numbers: %v \n Uris: %v", r.ClassNumbers, r.Uris)
}

type ResponseForUpdateClassSchema SchemaResponse

func (r *ResponseForUpdateClassSchema) GetRawMessage() []byte {
	return r.RawMessage
}

func (r *ResponseForUpdateClassSchema) SetRawMessage(raw []byte) {
	r.RawMessage = raw
}
