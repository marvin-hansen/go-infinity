package rest

import (
	"go-infinity/rest/index"
)

// AddIndex Creates an index for the given class using the given attributes as keys.
func (c *Client) AddIndex(req *index.RequestForAddIndex) (*index.ResponseForAddIndex, error) {
	results := index.NewResponseForAddIndex()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// DeleteIndex Deletes the named index from the federated database.
func (c *Client) DeleteIndex(req *index.RequestForDeleteIndex) error {
	if err := c.requestWithoutReturnValue(req); err != nil {
		return err
	}
	return nil
}
