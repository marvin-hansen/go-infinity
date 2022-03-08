package query

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/query"
	"go-infinity/test/shared"
	"testing"
)

func TestQuery(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	q1 := "FROM FleetData.Vehicle RETURN *;"

	res, err := c.Query(query.NewRequestForQuery(q1))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	println(string(res.GetRawMessage()))

}
