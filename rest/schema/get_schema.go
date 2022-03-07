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

type ResponseForRetrieveSchema []Class

type Class struct {
	ClassName       string       `json:"className"`
	ClassNumber     string       `json:"classNumber"`
	ShapeNumber     string       `json:"shapeNumber"`
	IsReferenceable bool         `json:"isReferenceable"`
	IsInternal      bool         `json:"isInternal"`
	IsDeleted       bool         `json:"isDeleted"`
	SuperClass      interface{}  `json:"superClass"`
	Attributes      []*Attribute `json:"attributes"`
}

func (s Class) String() string {
	return fmt.Sprintf("[Class]: \n ClassName: %v \n  ClassNumber: %v \n ShapeNumber: %v \n SuperClass: %v \n  IsReferenceable: %v \n  IsInternal: %v \n  IsDeleted: %v \n  Attributes: %v ",
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

type Attribute struct {
	AttributeName        string                `json:"attributeName"`
	LogicalType          string                `json:"logicalType"`
	Encoding             string                `json:"encoding"`
	Storage              string                `json:"storage"`
	ReferencedClass      string                `json:"referencedClass,omitempty"`
	InverseAttribute     string                `json:"inverseAttribute,omitempty"`
	CollectionTypeName   string                `json:"collectionTypeName,omitempty"`
	KeySpecification     *KeySpecification     `json:"keySpecification,omitempty"`
	ElementSpecification *ElementSpecification `json:"elementSpecification,omitempty"`
}

func (s *Attribute) String() string {
	if s == nil {
		return ""
	} else {
		return fmt.Sprintf("[\n  AttributeName: %v \n  LogicalType: %v \n  Encoding: %v \n  Storage: %v \n  ReferencedClass: %v \n  InverseAttribute: %v \n  CollectionTypeName: %v \n  KeySpecification: %v \n  ElementSpecification: %v \n",
			s.AttributeName,
			s.LogicalType,
			s.Encoding,
			s.Storage,
			s.ReferencedClass,
			s.InverseAttribute,
			s.CollectionTypeName,
			s.KeySpecification,
			s.ElementSpecification,
		)
	}
}

type ElementSpecification struct {
	LogicalType      string `json:"logicalType,omitempty"`
	Encoding         string `json:"encoding,omitempty"`
	ReferencedClass  string `json:"referencedClass,omitempty"`
	InverseAttribute string `json:"inverseAttribute,omitempty"`
}

func (s *ElementSpecification) String() string {
	if s == nil {
		return ""
	} else {
		return fmt.Sprintf("\n LogicalType: %v \n Encoding: %v \n ReferencedClass: %v \n InverseAttribute: %v \n ",
			s.LogicalType,
			s.Encoding,
			s.ReferencedClass,
			s.InverseAttribute,
		)
	}
}

type KeySpecification struct {
	LogicalType string `json:"logicalType,omitempty"`
	Encoding    string `json:"encoding,omitempty"`
	Storage     string `json:"storage,omitempty"`
}

func (s *KeySpecification) String() string {
	if s == nil {
		return ""
	} else {
		return fmt.Sprintf("\n LogicalType: %v \n Encoding: %v \n Storage: %v \n ",
			s.LogicalType,
			s.Encoding,
			s.Storage,
		)
	}
}
