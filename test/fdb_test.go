package test

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/fdb"
	"testing"
)

func TestRequestForBootPath(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, clientError)

	res, err := c.DBPATH(fdb.NewRequestForBootPath())
	assert.NoError(t, err)
	assert.NotNil(t, res)

	println(res.String())
}
