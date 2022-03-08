package query

import (
	"fmt"
	"go-infinity/rest/flags"
	"net/http"
)

const defaultVersion = flags.VERSION

//**// Request //**//

func getPayloadQuery(query string) (payload []byte) {

	str := fmt.Sprintf(`{"query": "%v", "language": "do" }`, query)

	payload = []byte(str)
	return payload
}

func NewRequestForQuery(query string) *RequestForQuery {
	payload := getPayloadQuery(query)
	return &RequestForQuery{payload: payload, version: defaultVersion}
}

type RequestForQuery struct {
	payload []byte
	version string
}

func (req *RequestForQuery) Path() string {
	return "/" + req.version + "/query"
}

func (req *RequestForQuery) Method() string {
	return http.MethodPost
}

func (req *RequestForQuery) Query() string {
	return ""
}

func (req *RequestForQuery) Payload() []byte {
	return req.payload
}

func (req *RequestForQuery) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForQuery() *ResponseForQuery {
	return new(ResponseForQuery)
}

func (r *ResponseForQuery) GetRawMessage() (raw []byte) {
	return r.RawMessage
}

func (r *ResponseForQuery) SetRawMessage(raw []byte) {
	r.RawMessage = raw
}

type ResponseForQuery struct {
	Entries []struct {
		Class string `json:"__class__"`
	} `json:"entries"`
	RawMessage []byte
}

type Entry struct {
	ClassName string `json:"__class__"`
}

func (r *Entry) String() string {
	return fmt.Sprintf("ClassName: %v \n",
		r.ClassName,
	)
}
