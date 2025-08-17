package builtin

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExit(t *testing.T) {
	old := exitFn
	defer func() {
		exitFn = old
	}()

	called, got := false, -1
	exitFn = func(code int) {
		called = true
		got = code
	}

	// Too many arguments
	msg, err := Exit("exit", []string{"0", "0"})
	require.Error(t, err)
	require.Equal(t, "", msg)

	// Invalid code
	msg, err = Exit("exit", []string{"-1"})
	require.Error(t, err)
	require.Equal(t, "", msg)

	// Invalid code
	msg, err = Exit("exit", []string{"2324"})
	require.Error(t, err)
	require.Equal(t, "", msg)

	// Valid code
	msg, err = Exit("exit", []string{"0"})
	require.NoError(t, err)
	require.Equal(t, "\n", msg)
	require.True(t, called)
	require.Equal(t, 0, got)
}
