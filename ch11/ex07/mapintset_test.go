package intset

import (
	"math/rand"
	"testing"
)

func BenchmarkMapIntSetAdd(b *testing.B) {
	rand.Seed(0)

	for i := 0; i < b.N; i++ {
		mis := &MapIntSet{}
		for j := 0; j < loopn; j++ {
			mis.Add(rand.Intn(maxn))
		}
	}
}

func BenchmarkMapIntSetHas(b *testing.B) {
	rand.Seed(0)
	mis := &MapIntSet{}
	for i := 0; i < loopn; i++ {
		mis.Add(rand.Intn(maxn))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mis.Has(rand.Intn(maxn))
	}
}

func BenchmarkMapIntSetString(b *testing.B) {
	rand.Seed(0)
	mis := &MapIntSet{}
	for i := 0; i < loopn; i++ {
		mis.Add(rand.Intn(maxn))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mis.String()
	}
}
