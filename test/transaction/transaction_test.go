package transaction

import (
	"github.com/stretchr/testify/assert"
	"go-infinity/rest"
	"go-infinity/rest/transaction"
	"go-infinity/test/shared"
	"testing"
)

// Throws: {"message":"While executing operation 1, encountered: org.restlet.representation.StringRepresentation cannot be cast to org.restlet.ext.jackson.JacksonRepresentation","type":"SystemError"}

func TestTransaction(t *testing.T) {
	c := rest.NewClient(nil)
	assert.NotNil(t, c, shared.ClientError)

	tx := testTransaction
	res, err := c.Transaction(transaction.NewRequestForTransaction(tx))
	assert.NoError(t, err)
	assert.NotNil(t, res)

}
