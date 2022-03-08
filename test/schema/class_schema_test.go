package schema

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/schema"
	"go-infinity/test/shared"
	"testing"
)

func TestGetClassSchema(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	classID := "FleetData.Customer"
	res, err := c.GetClassSchema(schema.NewRequestForGetClassSchema(classID))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	println(res.String())

}

func TestDeleteClassSchema(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)
	classID := "FleetData.Customer"
	err := c.DeleteClassSchema(schema.NewRequestForDeleteClassSchema(classID))
	assert.NoError(t, err)
}
