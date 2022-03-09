package transaction

import (
	"go-infinity/rest/flags"
	"net/http"
)

const defaultVersion = flags.VERSION

//**// Request //**//

func NewRequestForTransaction(transaction string) *RequestForTransaction {
	return &RequestForTransaction{payload: []byte(transaction), version: defaultVersion}
}

type RequestForTransaction struct {
	payload []byte
	version string
}

func (req *RequestForTransaction) Path() string {
	return "/" + req.version + "/transaction"
}

func (req *RequestForTransaction) Method() string {
	return http.MethodPost
}

func (req *RequestForTransaction) Query() string {
	return ""
}

func (req *RequestForTransaction) Payload() []byte {
	return req.payload
}

func (req *RequestForTransaction) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForTransaction() *ResponseForTransaction {
	return new(ResponseForTransaction)
}

type ResponseForTransaction []Response

type Response struct {
	ResponseCode string `json:"responseCode"`
	Result       Result `json:"result"`
}

type Result struct {
	Uri string `json:"uri"`
}

func (r *ResponseForTransaction) GetRawMessage() (raw []byte) {
	return raw
}

func (r *ResponseForTransaction) SetRawMessage(raw []byte) {
	//r.RawMessage = raw
}
