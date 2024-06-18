package breaker

import (
	"sync/atomic"
	"testing"
)

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
	breaker := atomic.Bool{}

	b.ResetTimer()

	for range b.N {
		if breaker.Load() {
			return
		}
	}
}
