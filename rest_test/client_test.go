package rest_test

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"testing"
)

func TestNewClient(t *testing.T) {

	c := rest.NewClient(nil)
	assert.NotNil(t, c, "Should not be nil")

}
