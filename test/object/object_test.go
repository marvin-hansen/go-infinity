package object

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/schema"
	"go-infinity/test/shared"
	"testing"
)

// TestPre adds a schema to the DB so that objects can be created, modified, and deleted
func TestPre(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	testSchema := []byte(shared.TestBigSchema)
	res, err := c.AddSchema(schema.NewRequestForAddSchema(testSchema))
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

// TestPost Deletes all test data and the schema to clean everything up after testing.
func TestPost(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	classID := "FleetData.Customer"
	err := c.DeleteClassSchema(schema.NewRequestForDeleteClassSchema(classID))
	assert.NoError(t, err)
}
