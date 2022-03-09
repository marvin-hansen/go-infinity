package object

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/object"
	"go-infinity/rest/schema"
	"go-infinity/test/shared"
	"testing"
)

const objId = "3-3-1-8"

// TestPrepare adds a schema to the DB so that objects can be created, modified, and deleted
func TestPrepare(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	nameSpace := "FleetData"
	_, errNameSpace := c.GetNameSpaceSchema(schema.NewRequestForGetNameSpace(nameSpace))
	if errNameSpace == nil {
		println("No Schema, create one!")
		testSchema := []byte(shared.TestBigSchema)
		res, err := c.AddSchema(schema.NewRequestForAddSchema(testSchema))
		assert.NoError(t, err)
		assert.NotNil(t, res)
	}
}

func TestAddObject(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	testObject := []byte(shared.TestAddObjectVehicle)
	res, err := c.AddObject(object.NewRequestForAddObject(testObject))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	println(res.String())
}

func TestUpdateObject(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	objectId := objId // "3-3-1-4"
	testUpdate := []byte(shared.TestUpdateObjectVehicle)
	errUpdate := c.UpdateObject(object.NewRequestForUpdateObject(objectId, testUpdate))
	assert.NoError(t, errUpdate)

	res, err := c.GetObject(object.NewRequestForGetObject(objectId))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	println(res.String())
}

func TestGetObject(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	objectId := objId //"3-3-1-4"
	res, err := c.GetObject(object.NewRequestForGetObject(objectId))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	println(res.String())

}

func TestDeleteObject(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	objectId := objId // "3-3-1-4"
	err := c.DeleteObject(object.NewRequestForDeleteObject(objectId))
	assert.NoError(t, err)
}

// TestTeardown Deletes all test data and the schema to clean everything up after testing.
func TestTeardown(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	nameSpace := "FleetData"
	namespaceSchema, errNameSpace := c.GetNameSpaceSchema(schema.NewRequestForGetNameSpace(nameSpace))
	assert.NoError(t, errNameSpace)

	if namespaceSchema != nil {
		println("Delete test Schema")
		// can't delete complex schema; must delete & re-create entire DB....
	}
}
