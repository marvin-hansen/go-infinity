package tool

import (
	"fmt"
	"net/http"
)

//**// Request //**//

func NewRequestForGetTool() *RequestForGetTool {
	return &RequestForGetTool{defaultVersion}
}

type RequestForGetTool struct {
	version string
}

func (req *RequestForGetTool) Path() string {
	return "/" + req.version + "/tool"
}

func (req *RequestForGetTool) Method() string {
	return http.MethodGet
}

func (req *RequestForGetTool) Query() string {
	return ""
}

func (req *RequestForGetTool) Payload() []byte {
	return nil
}

func (req *RequestForGetTool) ResponseCode() int {
	return 200 // ok
}

//**// Response //**//

func NewResponseForGetTool() *ResponseForGetTool {
	return new(ResponseForGetTool)
}

type ResponseForGetTool []Tool

type Tool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r Tool) String() string {
	return fmt.Sprintf("Name: %v, \n Description: %v \n",
		r.Name,
		r.Description,
	)
}

func (r *ResponseForGetTool) GetRawMessage() (raw []byte) { return raw }

func (r *ResponseForGetTool) SetRawMessage(raw []byte) {}
