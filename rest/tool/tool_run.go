package tool

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForRunTool() *RequestForRunTool {
	return &RequestForRunTool{defaultVersion}
}

type RequestForRunTool struct {
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
	return nil
}

func (req *RequestForRunTool) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForRunTool() *ResponseForRunTool {
	return new(ResponseForRunTool)
}

func (r ResponseForRunTool) String() string {
	return fmt.Sprintf("ReturnCode: %v \n ToolOutput: %v \n ToolError: %v \n",
		r.ReturnCode,
		r.ToolOutput,
		r.ToolError,
	)
}

type ResponseForRunTool Response

type Response struct {
	ReturnCode int
	ToolOutput string
	ToolError  string
	RawMessage []byte
}

func (r *ResponseForRunTool) GetRawMessage() []byte {
	return r.RawMessage
}

func (r *ResponseForRunTool) SetRawMessage(raw []byte) {
	r.RawMessage = raw
}
