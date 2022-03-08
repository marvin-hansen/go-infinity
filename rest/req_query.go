package rest

import "go-infinity/rest/query"

func (c *Client) Query(req *query.RequestForQuery) (*query.ResponseForQuery, error) {
	results := query.NewResponseForQuery()
	if err := c.requestQuery(req, results); err != nil {
		return nil, err
	}
	return results, nil
}
