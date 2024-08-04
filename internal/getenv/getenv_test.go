package getenv

import (
	"testing"

	"github.com/akramarenkov/goresearch/internal/consts"

	"github.com/stretchr/testify/require"
)

func TestSize(t *testing.T) {
	size, err := Size(1024)
	require.NoError(t, err)
	require.Equal(t, 1024, size)

	t.Setenv(consts.EnvPrefix+sizeKey, "1")

	size, err = Size(1024)
	require.NoError(t, err)
	require.Equal(t, 1, size)

	t.Setenv(consts.EnvPrefix+sizeKey, "1.")

	_, err = Size(1024)
	require.Error(t, err)
}
