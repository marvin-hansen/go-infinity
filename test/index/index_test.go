package index

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/index"
	"go-infinity/test/shared"
	"testing"
)

func TestAddIndex(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	idx := testAddIndex
	res, err := c.AddIndex(index.NewRequestForAddIndex(idx))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	println(res.String())
}

func TestDeleteIndex(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	idxID := testDeleteIndex
	err := c.DeleteIndex(index.NewRequestForDeleteIndex(idxID))
	assert.NoError(t, err)
}
