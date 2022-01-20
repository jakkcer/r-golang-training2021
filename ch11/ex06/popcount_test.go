package popcount

import (
	"math/rand"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(rand.Uint64())
	}
}

func BenchmarkPopCount64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount64(rand.Uint64())
	}
}

func BenchmarkClearPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClearPopCount(rand.Uint64())
	}
}
