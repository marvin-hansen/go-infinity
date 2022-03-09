package rest

import (
	"go-infinity/rest/index"
)

func (c *Client) AddIndex(req *index.RequestForAddIndex) (*index.ResponseForAddIndex, error) {
	results := index.NewResponseForAddIndex()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (c *Client) DeleteIndex(req *index.RequestForDeleteIndex) error {
	if err := c.requestWithoutReturnValue(req); err != nil {
		return err
	}
	return nil
}
