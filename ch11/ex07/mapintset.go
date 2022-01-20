package intset

import (
	"sort"
	"strconv"
)

type MapIntSet struct {
	words map[int]bool
}

func (s *MapIntSet) Has(x int) bool {
	if s.words == nil {
		return false
	}
	return s.words[x]
}

func (s *MapIntSet) Add(x int) {
	if s.words == nil {
		s.words = make(map[int]bool)
	}
	s.words[x] = true
}

func (s *MapIntSet) UnionWith(t *MapIntSet) {
	if t.words != nil && s.words == nil {
		s.words = make(map[int]bool)
	}
	for x := range t.words {
		s.words[x] = true
	}
}

func (s *MapIntSet) String() string {
	var sw []int
	for x := range s.words {
		sw = append(sw, x)
	}
	sort.Ints(sw)
	b := []byte{'{'}
	for i, word := range sw {
		if i > 0 {
			b = append(b, ' ')
		}
		b = strconv.AppendInt(b, int64(word), 10)
	}
	b = append(b, '}')
	return string(b)
}
