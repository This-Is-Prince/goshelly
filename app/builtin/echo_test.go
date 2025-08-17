package builtin

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEcho(t *testing.T) {
	msg, err := Echo("echo", []string{"hi", "Hey"})
	require.NoError(t, err)
	require.Equal(t, "hi Hey\n", msg)
}
