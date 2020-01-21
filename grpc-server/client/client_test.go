package client

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGRPCClient(t *testing.T) {
	n, err := GRPCClient("0.0.0.0:5101")
	require.NoError(t, err, "Unexpected error")
	assert.Equal(t, 10, n)
}
