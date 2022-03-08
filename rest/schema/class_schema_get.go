package schema

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetClassSchema(classID string) *RequestForGetClassSchema {
	return &RequestForGetClassSchema{classID: classID, version: defaultVersion}
}

// RequestForGetClassSchema Retrieves the schema representation for a class identified by class name or class number.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestSchemaTypeGET.html%23
type RequestForGetClassSchema struct {
	classID string
	version string
}

func (req *RequestForGetClassSchema) Path() string {
	return "/" + req.version + "/" + "schema" + "/" + req.classID
}

func (req *RequestForGetClassSchema) Method() string {
	return http.MethodGet
}

func (req *RequestForGetClassSchema) Query() string {
	return ""
}

func (req *RequestForGetClassSchema) Payload() []byte {
	return nil
}

func (req *RequestForGetClassSchema) ResponseCode() int {
	return 200
}

//**// Response //**//

func NewResponseForGetClassSchema() *ResponseForGetClassSchema {
	return new(ResponseForGetClassSchema)
}

type ResponseForGetClassSchema Schema

func (s *ResponseForGetClassSchema) String() string {
	return fmt.Sprintf("[Schema]: \n ClassName: %v \n  ClassNumber: %v \n ShapeNumber: %v \n SuperClass: %v \n  IsReferenceable: %v \n  IsInternal: %v \n  IsDeleted: %v \n  Attributes: %v ",
		s.ClassName,
		s.ClassNumber,
		s.ShapeNumber,
		s.SuperClass,
		s.IsReferenceable,
		s.IsInternal,
		s.IsDeleted,
		s.Attributes,
	)
}

func (s *ResponseForGetClassSchema) GetRawMessage() []byte {
	return s.RawMessage
}

func (s *ResponseForGetClassSchema) SetRawMessage(raw []byte) {
	s.RawMessage = raw
}
