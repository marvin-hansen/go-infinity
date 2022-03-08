package server

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/server"
	"go-infinity/test/shared"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)
}

func TestRequestForResources(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	res, err := c.Resources(server.NewRequestForResources("v1"))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	// println(string(res.GetRawMessage()))
}

func TestRequestForUptime(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	res, err := c.Uptime(server.NewRequestForUptime())
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
