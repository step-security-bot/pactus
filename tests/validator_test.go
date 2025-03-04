package tests

import (
	"testing"

	"github.com/pactus-project/pactus/crypto"
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getValidator(_ *testing.T, addr crypto.Address) *pactus.ValidatorInfo {
	res, err := tBlockchain.GetValidator(tCtx,
		&pactus.GetValidatorRequest{Address: addr.String()})
	if err != nil {
		return nil
	}
	return res.Validator
}

func TestGetValidator(t *testing.T) {
	val := getValidator(t, tSigners[tNodeIdx2][0].Address())
	require.NotNil(t, val)
	assert.Equal(t, val.Number, int32(1))
}
