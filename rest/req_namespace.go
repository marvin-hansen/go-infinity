package rest

import "go-infinity/rest/schema"

// GetAllNameSpaces RequestForGetAllNameSpaces Returns a list of all schema namespaces in the federated database.
func (c *Client) GetAllNameSpaces(req *schema.RequestForGetAllNameSpaces) (*schema.ResponseForGetAllNameSpaces, error) {
	results := schema.NewResponseForGetAllNameSpaces()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// GetNameSpaceSchema Returns the schema representations for all classes in the specified namespace.
func (c *Client) GetNameSpaceSchema(req *schema.RequestForGetNameSpace) (*schema.ResponseForGetNameSpace, error) {
	results := schema.NewResponseForGetNameSpace()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}
