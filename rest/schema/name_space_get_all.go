package schema

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetAllNameSpaces() *RequestForGetAllNameSpaces {
	return &RequestForGetAllNameSpaces{defaultVersion}
}

// RequestForGetAllNameSpaces Returns a list of all schema namespaces in the federated database.
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestSchemaNamespaceGet.html%23
type RequestForGetAllNameSpaces struct {
	version string
}

func (req *RequestForGetAllNameSpaces) Path() string {
	return "/" + req.version + "/" + "schema" + "/" + "namespace"
}

func (req *RequestForGetAllNameSpaces) Method() string {
	return http.MethodGet
}

func (req *RequestForGetAllNameSpaces) Query() string {
	return ""
}

func (req *RequestForGetAllNameSpaces) Payload() []byte {
	return nil
}

func (req *RequestForGetAllNameSpaces) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetAllNameSpaces() *ResponseForGetAllNameSpaces {
	return new(ResponseForGetAllNameSpaces)
}

type ResponseForGetAllNameSpaces []NameSpaceNames

func (r *ResponseForGetAllNameSpaces) GetRawMessage() (raw []byte) { return raw }

func (r *ResponseForGetAllNameSpaces) SetRawMessage(raw []byte) {}

type NameSpaceNames struct {
	Name string `json:"name"`
}

func (r *NameSpaceNames) String() string {
	if r == nil {
		return ""
	} else {
		return fmt.Sprintf("\n Namespace: %v ",
			r.Name,
		)
	}
}
