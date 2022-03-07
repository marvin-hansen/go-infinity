package rest

import "go-infinity/rest/schema"

// AddSchema Adds one or more new class descriptions to the schema of a federated database.
func (c *Client) AddSchema(req *schema.RequestForAddSchema) (*schema.ResponseForAddSchema, error) {
	results := schema.NewResponseForAddSchema()
	if err := c.request(req, results, http201); err != nil {
		return nil, err
	}
	return results, nil
}

// UpdateSchema Provides one or more class descriptions, each of which results in the creation of a new class or the addition of new attributes to an existing class.
func (c *Client) UpdateSchema(req *schema.RequestForUpdateSchema) (*schema.ResponseForUpdateSchema, error) {
	results := schema.NewResponseForUpdateSchema()
	if err := c.request(req, results, http201); err != nil {
		return nil, err
	}
	return results, nil
}

// GetSchema Retrieves a representation of the complete schema for a federated database.
func (c *Client) GetSchema(req *schema.RequestForRetrieveSchema) (*schema.ResponseForRetrieveSchema, error) {
	results := schema.NewResponseForRetrieveSchema()
	if err := c.request(req, results, http200); err != nil {
		return nil, err
	}
	return results, nil
}
