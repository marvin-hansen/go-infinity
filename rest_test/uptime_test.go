package rest

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/server"
	"testing"
)

func TestRequestForUptime(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, "Should not be nil")

	res, err := c.Uptime(server.NewRequestForUptime())
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
