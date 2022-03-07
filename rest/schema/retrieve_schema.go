package schema

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForRetrieveSchema() *RequestForRetrieveSchema {
	return &RequestForRetrieveSchema{defaultVersion}
}

// RequestForRetrieveSchema Retrieves a representation of the complete schema for a federated database.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestSchemaGET.html%23
type RequestForRetrieveSchema struct {
	version string
}

func (req *RequestForRetrieveSchema) Path() string {
	return "/" + req.version + "/schema"
}

func (req *RequestForRetrieveSchema) Method() string {
	return http.MethodGet
}

func (req *RequestForRetrieveSchema) Query() string {
	return ""
}

func (req *RequestForRetrieveSchema) Payload() []byte {
	return nil
}

//**// Response //**//

func NewResponseForRetrieveSchema() *ResponseForRetrieveSchema {
	return new(ResponseForRetrieveSchema)
}

func (r ResponseForRetrieveSchema) String() string {
	return fmt.Sprintf("[Schema]: %v \n", r.Classes)
}

type ResponseForRetrieveSchema Schema

type Schema struct {
	Classes []Class
}

type Class struct {
	ClassName       string       `json:"className"`
	ClassNumber     string       `json:"classNumber"`
	ShapeNumber     string       `json:"shapeNumber"`
	SuperClass      string       `json:"superClass"`
	IsReferenceable bool         `json:"isReferenceable"`
	IsInternal      bool         `json:"isInternal"`
	IsDeleted       bool         `json:"isDeleted"`
	Attributes      []*Attribute `json:"attributes"`
}

func (c Class) String() string {
	return fmt.Sprintf("[Class]: \n ClassName: %v \n  ClassNumber: %v \n ShapeNumber: %v \n SuperClass: %v \n  IsReferenceable: %v \n  IsInternal: %v \n  IsDeleted: %v \n  Attributes: %v ",
		c.ClassName,
		c.ClassNumber,
		c.ShapeNumber,
		c.SuperClass,
		c.IsReferenceable,
		c.IsInternal,
		c.IsDeleted,
		c.Attributes,
	)
}

type Attribute struct {
	AttributeName string `json:"attributeName"`
	LogicalType   string `json:"logicalType"`
	Encoding      string `json:"encoding"`
	Storage       string `json:"storage"`
}

func (a Attribute) String() string {
	return fmt.Sprintf("AttributeName: %v \n  LogicalType: %v \n Encoding: %v \n Storage: %v ",
		a.AttributeName,
		a.LogicalType,
		a.Encoding,
		a.Storage,
	)
}
