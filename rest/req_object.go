package rest

import (
	"go-infinity/rest/object"
)

// AddObject Adds a new object to a federated database and assigns an OID.
func (c *Client) AddObject(req *object.RequestForAddObject) (*object.ResponseForAddObject, error) {
	results := object.NewResponseForAddObject()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// UpdateObject Modifies an existing object by changing attribute values.
func (c *Client) UpdateObject(req *object.RequestForUpdateObject) (*object.ResponseForUpdateObject, error) {
	results := object.NewResponseForUpdateObject()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// GetObject Retrieves a representation of the object with the given OID.
func (c *Client) GetObject(req *object.RequestForGetObject) (*object.ResponseForGetObject, error) {
	results := object.NewResponseForGetObject()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// DeleteObject Deletes the object with the given OID from the federated database.
func (c *Client) DeleteObject(req *object.RequestForDeleteObject) error {
	if err := c.requestWithoutReturnValue(req); err != nil {
		return err
	}
	return nil
}
