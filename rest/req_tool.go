package rest

import (
	"go-infinity/rest/tool"
)

func (c *Client) GetAllTools(req *tool.RequestForGetTool) (*tool.ResponseForGetTool, error) {
	results := tool.NewResponseForGetTool()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (c *Client) RunTool(req *tool.RequestForRunTool) (*tool.ResponseForRunTool, error) {
	results := tool.NewResponseForRunTool()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}
