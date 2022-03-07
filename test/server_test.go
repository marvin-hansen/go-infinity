package test

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/server"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, clientError)
}

func TestRequestForResources(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, clientError)

	res, err := c.Resources(server.NewRequestForResources("v1"))
	assert.NoError(t, err)
	assert.NotNil(t, res)
	// println(res.String())
}

func TestRequestForUptime(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, clientError)

	res, err := c.Uptime(server.NewRequestForUptime())
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
