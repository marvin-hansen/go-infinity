package tool

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/tool"
	"go-infinity/test/shared"
	"testing"
)

func TestGetAllTools(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	res, err := c.GetAllTools(tool.NewRequestForGetTool())
	assert.NoError(t, err)
	assert.NotNil(t, res)

	for _, e := range *res {
		println(e.String())
	}
}
func TestGetRunTool(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	runCommand := testToolRun
	res, err := c.RunTool(tool.NewRequestForRunTool(runCommand))
	assert.NoError(t, err)
	assert.NotNil(t, res)

	println(res.String())

}
