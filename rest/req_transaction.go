package rest

import (
	"go-infinity/rest/transaction"
)

func (c *Client) Transaction(req *transaction.RequestForTransaction) (*transaction.ResponseForTransaction, error) {
	results := transaction.NewResponseForTransaction()
	if err := c.requestQuery(req, results); err != nil {
		return nil, err
	}
	return results, nil
}
