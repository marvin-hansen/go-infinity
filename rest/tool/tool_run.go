package tool

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForRunTool(runCommandJson string) *RequestForRunTool {
	return &RequestForRunTool{payload: []byte(runCommandJson), version: defaultVersion}
}

type RequestForRunTool struct {
	payload []byte
	version string
}

func (req *RequestForRunTool) Path() string {
	return "/" + req.version + "/tool"
}

func (req *RequestForRunTool) Method() string {
	return http.MethodPost
}

func (req *RequestForRunTool) Query() string {
	return ""
}

func (req *RequestForRunTool) Payload() []byte {
	return req.payload
}

func (req *RequestForRunTool) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForRunTool() *ResponseForRunTool {
	return new(ResponseForRunTool)
}

func (r ResponseForRunTool) String() string {
	return fmt.Sprintf("ReturnCode: %v \nToolOutput: %v \nToolError: %v \n",
		r.ReturnCode,
		r.ToolOutput,
		r.ToolError,
	)
}

func (r *ResponseForRunTool) GetRawMessage() []byte {
	return r.RawMessage
}

func (r *ResponseForRunTool) SetRawMessage(raw []byte) {
	r.RawMessage = raw
}

type ResponseForRunTool Response

type Response struct {
	ReturnCode int
	ToolOutput string
	ToolError  string
	RawMessage []byte
}
