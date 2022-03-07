package test

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/schema"
	"testing"
)

func TestAddSchema(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, clientError)

	res, err := c.AddSchema(schema.NewRequestForAddSchema([]byte(testAddSchema)))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	println(res.String())

}

func TestUpdateSchema(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, clientError)

	res, err := c.UpdateSchema(schema.NewRequestForUpdateSchema([]byte(testUpdateSchema)))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	println(res.String())
}

func TestGetSchema(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, clientError)

	res, err := c.GetSchema(schema.NewRequestForRetrieveSchema())
	assert.NoError(t, err)
	assert.NotNil(t, res)

	for _, e := range *res {
		println(e.String())
	}
}
