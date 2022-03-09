package index

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/test/shared"
	"testing"
)

func TestGetAllIndices(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

}
