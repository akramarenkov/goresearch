package breaker

import (
	"os"
	"strconv"
	"sync/atomic"
	"testing"

	"github.com/akramarenkov/goresearch/internal/consts"

	"github.com/stretchr/testify/require"
)

func getTerminateFromEnv(b *testing.B) bool {
	env := os.Getenv(consts.EnvPrefix + "TERMINATE")

	if env == "" {
		return false
	}

	terminate, err := strconv.ParseBool(env)
	require.NoError(b, err)

	return terminate
}

func BenchmarkIdle(b *testing.B) {
	for range b.N {
		_ = b.N
	}
}

func BenchmarkChannelOpened(b *testing.B) {
	breaker := make(chan struct{})

	b.ResetTimer()

	for range b.N {
		select {
		case <-breaker:
			return
		default:
		}
	}
}

func BenchmarkChannelClosed(b *testing.B) {
	breaker := make(chan struct{})

	close(breaker)

	b.ResetTimer()

	for range b.N {
		select {
		case <-breaker:
		default:
			return
		}
	}
}

func BenchmarkAtomic(b *testing.B) {
	breaker := &atomic.Bool{}

	breaker.Store(getTerminateFromEnv(b))

	b.ResetTimer()

	for range b.N {
		if breaker.Load() {
			return
		}
	}
}
