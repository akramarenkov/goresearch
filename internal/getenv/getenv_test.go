package getenv

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnvVar(t *testing.T) {
	require.Equal(t, "GORESEARCH_NAME", envVar("NAME"))
}

func TestSize(t *testing.T) {
	const def = 2048

	size, err := Size(def)
	require.NoError(t, err)
	require.Equal(t, def, size)

	t.Setenv(envVar(keySize), "1")

	size, err = Size(def)
	require.NoError(t, err)
	require.Equal(t, 1, size)

	t.Setenv(envVar(keySize), "1.")

	_, err = Size(def)
	require.Error(t, err)
}

func TestInt(t *testing.T) {
	const (
		def = 20
		key = "INT"
	)

	value, err := Int(key, def)
	require.NoError(t, err)
	require.Equal(t, def, value)

	t.Setenv(envVar(key), "1")

	value, err = Int(key, def)
	require.NoError(t, err)
	require.Equal(t, 1, value)

	t.Setenv(envVar(key), "1.")

	_, err = Int(key, def)
	require.Error(t, err)
}

func TestBoolean(t *testing.T) {
	const (
		def = true
		key = "BOOL"
	)

	value, err := Boolean(key, def)
	require.NoError(t, err)
	require.True(t, value)

	t.Setenv(envVar(key), "false")

	value, err = Boolean(key, def)
	require.NoError(t, err)
	require.False(t, value)

	t.Setenv(envVar(key), "flase")

	_, err = Boolean(key, def)
	require.Error(t, err)
}
