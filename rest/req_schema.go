package rest

import "go-infinity/rest/schema"

// AddSchema Adds one or more new class descriptions to the schema of a federated database.
func (c *Client) AddSchema(req *schema.RequestForAddSchema) (*schema.ResponseForAddSchema, error) {
	results := schema.NewResponseForAddSchema()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// UpdateSchema Provides one or more class descriptions, each of which results in the creation of a new class or the addition of new attributes to an existing class.
func (c *Client) UpdateSchema(req *schema.RequestForUpdateSchema) (*schema.ResponseForUpdateSchema, error) {
	results := schema.NewResponseForUpdateSchema()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// UpdateClassSchema Modifies an existing class, either renaming the class or deleting, renaming, or adding attributes.
func (c *Client) UpdateClassSchema(req *schema.RequestForUpdateClassSchema) (*schema.ResponseForUpdateClassSchema, error) {
	results := schema.NewResponseForUpdateClassSchema()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// GetSchema Retrieves a representation of the complete schema for a federated database.
func (c *Client) GetSchema(req *schema.RequestForGetSchema) (*schema.ResponseForRetrieveSchema, error) {
	results := schema.NewResponseForGetSchema()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// GetClassSchema Retrieves the schema representation for a class identified by class name or class number.
func (c *Client) GetClassSchema(req *schema.RequestForGetClassSchema) (*schema.ResponseForGetClassSchema, error) {
	results := schema.NewResponseForGetClassSchema()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// DeleteClassSchema Deletes the class description for an existing type specified by fully qualified class name or class number.
// Before deleting a class, remove all objects of that class from the federated database. Objects of the class become inaccessible to applications after the class is deleted.
// Also, attributes of other classes that reference objects of the deleted class are deleted.
func (c *Client) DeleteClassSchema(req *schema.RequestForDeleteClassSchema) (*schema.ResponseForDeleteClassSchema, error) {
	results := schema.NewResponseForDeleteClassSchema()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}
