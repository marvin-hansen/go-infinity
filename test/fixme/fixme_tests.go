package fixme

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/schema"
	"go-infinity/test/shared"
	"testing"
)

func TestUpdateClassSchema(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	classID := "FleetData.Customer"
	updateSchema := []byte(shared.TestUpdateClassSchema)

	res, err := c.UpdateClassSchema(schema.NewRequestForUpdateClassSchema(classID, updateSchema))
	assert.NoError(t, err)
	assert.NotNil(t, res)
	println(res.String())
}
