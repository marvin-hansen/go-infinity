package schema

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetNameSpace(nameSpace string) *RequestForGetNameSpace {
	return &RequestForGetNameSpace{nameSpace: nameSpace, version: defaultVersion}
}

type RequestForGetNameSpace struct {
	nameSpace string
	version   string
}

func (req *RequestForGetNameSpace) Path() string {
	return "/" + req.version + "/" + "schema" + "/" + "namespace" + "/" + req.nameSpace
}

func (req *RequestForGetNameSpace) Method() string {
	return http.MethodGet
}

func (req *RequestForGetNameSpace) Query() string {
	return ""
}

func (req *RequestForGetNameSpace) Payload() []byte {
	return nil
}

func (req *RequestForGetNameSpace) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetNameSpace() *ResponseForGetNameSpace {
	return new(ResponseForGetNameSpace)
}

type ResponseForGetNameSpace Namespace

func (s *ResponseForGetNameSpace) GetRawMessage() []byte {
	return s.RawMessage
}

func (s *ResponseForGetNameSpace) SetRawMessage(raw []byte) {
	s.RawMessage = raw
}

func (s *ResponseForGetNameSpace) String() string {
	return fmt.Sprintf("Name: %v \n Classes: %v \n  ",
		s.Name,
		s.Classes,
	)
}
