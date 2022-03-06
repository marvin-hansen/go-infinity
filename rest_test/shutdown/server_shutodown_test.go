package shutdown

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/server"
	"testing"
)

func TestRequestForShutdown(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, "Client should not be nil")

	// server shutdown req doesn't return anything, obviously.
	err := c.ShutdownServer(server.NewRequestForShutdown())
	assert.NoError(t, err)
}
