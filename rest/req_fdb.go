package rest

import "go-infinity/rest/fdb"

func (c *Client) DBPATH(req *fdb.RequestForBootPath) (*fdb.ResponseForBootPath, error) {
	results := fdb.NewResponseForBootPath()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}
