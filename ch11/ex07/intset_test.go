package intset

import (
	"math/rand"
	"testing"
)

const (
	maxn  = 1000000
	loopn = 10000
)

func BenchmarkIntSetAdd(b *testing.B) {
	rand.Seed(0)

	for i := 0; i < b.N; i++ {
		is := &IntSet{}
		for j := 0; j < loopn; j++ {
			is.Add(rand.Intn(maxn))
		}
	}
}

func BenchmarkIntSetHas(b *testing.B) {
	rand.Seed(0)
	is := &IntSet{}
	for i := 0; i < loopn; i++ {
		is.Add(rand.Intn(maxn))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		is.Has(rand.Intn(maxn))
	}
}

func BenchmarkIntSetString(b *testing.B) {
	rand.Seed(0)
	is := &IntSet{}
	for i := 0; i < loopn; i++ {
		is.Add(rand.Intn(maxn))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		is.String()
	}
}
