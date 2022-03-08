package namespace

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/schema"
	"go-infinity/test/shared"
	"testing"
)

func TestGetAllNameSpaces(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	res, err := c.GetAllNameSpaces(schema.NewRequestForGetAllNameSpaces())
	assert.NoError(t, err)
	assert.NotNil(t, res)

	for _, e := range *res {
		println(e.String())
	}
}

func TestGetNameSpaces(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	nameSpace := "FleetData"
	res, err := c.GetNameSpaceSchema(schema.NewRequestForGetNameSpace(nameSpace))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	println(res.String())
}
