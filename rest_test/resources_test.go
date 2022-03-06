package rest

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/server"
	"testing"
)

func TestRequestForResources(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, "Should not be nil")

	res, err := c.Resources(server.NewRequestForResources("v1"))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	// println(res.String())
}
