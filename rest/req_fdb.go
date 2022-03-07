package rest

import "go-infinity/rest/fdb"

// DBPATH Retrieves the path to the boot file for the current federated database.
func (c *Client) DBPATH(req *fdb.RequestForBootPath) (*fdb.ResponseForBootPath, error) {
	results := fdb.NewResponseForBootPath()

	if err := c.request(req, results, http200); err != nil {
		return nil, err
	}
	return results, nil
}
