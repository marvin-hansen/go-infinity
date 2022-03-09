package index

import (
	"net/http"
)

//**// Request //**//

func NewRequestForDeleteIndex(indexId string) *RequestForDeleteIndex {
	return &RequestForDeleteIndex{IndexId: indexId, version: defaultVersion}
}

type RequestForDeleteIndex struct {
	IndexId string
	version string
}

func (req *RequestForDeleteIndex) Path() string {
	return "/" + req.version + "/index" + "/" + req.IndexId
}

func (req *RequestForDeleteIndex) Method() string {
	return http.MethodDelete
}

func (req *RequestForDeleteIndex) Query() string {
	return ""
}

func (req *RequestForDeleteIndex) Payload() []byte {
	return nil
}

func (req *RequestForDeleteIndex) ResponseCode() int {
	return 204 // No Content
}

//**// Response //**//

// No response due to 204 return code.
