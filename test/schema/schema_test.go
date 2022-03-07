package schema

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/schema"
	"go-infinity/test/shared"
	"testing"
)

func TestAddSchema(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	res, err := c.AddSchema(schema.NewRequestForAddSchema([]byte(shared.TestAddSchema)))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	println(res.String())

}

func TestUpdateSchema(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	res, err := c.UpdateSchema(schema.NewRequestForUpdateSchema([]byte(shared.TestUpdateSchema)))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	println(res.String())
}

func TestGetSchema(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	res, err := c.GetSchema(schema.NewRequestForGetSchema())
	assert.NoError(t, err)
	assert.NotNil(t, res)

	for _, e := range *res {
		println(e.String())
	}
}
