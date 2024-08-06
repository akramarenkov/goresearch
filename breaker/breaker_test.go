package breaker

import (
	"sync/atomic"
	"testing"

	"github.com/akramarenkov/goresearch/internal/getenv"

	"github.com/stretchr/testify/require"
)

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

	terminate, err := getenv.Terminate(false)
	require.NoError(b, err)

	breaker.Store(terminate)

	b.ResetTimer()

	for range b.N {
		if breaker.Load() {
			return
		}
	}
}
