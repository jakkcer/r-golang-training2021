package intset

import (
	"testing"
)

func newIntSet(a ...int) *IntSet {
	intSet := &IntSet{}
	for _, x := range a {
		intSet.Add(x)
	}
	return intSet
}

func TestIntSetHas(t *testing.T) {
	tests := []struct {
		intSet   *IntSet
		expected map[int]bool
	}{
		{
			intSet: newIntSet(1, 2, 3, 10),
			expected: map[int]bool{
				0:  false,
				1:  true,
				2:  true,
				3:  true,
				4:  false,
				5:  false,
				9:  false,
				10: true,
				11: false,
			},
		},
	}
	for _, test := range tests {
		for x, expected := range test.expected {
			if got := test.intSet.Has(x); got != expected {
				t.Errorf("Has(%d) expected: %v, but got: %v", x, expected, got)
			}
		}
	}
}

func TestIntSetAdd(t *testing.T) {
	tests := []struct {
		intSet   *IntSet
		adds     []int
		expected map[int]bool
	}{
		{
			intSet: newIntSet(1, 2, 3, 10),
			adds:   []int{0, 5, 100},
			expected: map[int]bool{
				0:   true,
				1:   true,
				2:   true,
				3:   true,
				4:   false,
				5:   true,
				9:   false,
				10:  true,
				11:  false,
				99:  false,
				100: true,
				101: false,
			},
		},
	}

	for _, test := range tests {
		for _, x := range test.adds {
			test.intSet.Add(x)
		}
		for x, expected := range test.expected {
			if got := test.intSet.Has(x); got != expected {
				t.Errorf("Has(%d) expected: %v, but got: %v", x, expected, got)
			}
		}
	}
}

func TestIntSetUnionWith(t *testing.T) {
	tests := []struct {
		a, b     *IntSet
		expected map[int]bool
	}{
		{
			a: newIntSet(1, 2, 3, 10),
			b: newIntSet(50, 55),
			expected: map[int]bool{
				0:  false,
				1:  true,
				2:  true,
				3:  true,
				4:  false,
				9:  false,
				10: true,
				11: false,
				49: false,
				50: true,
				51: false,
				54: false,
				55: true,
				56: false,
			},
		},
		{
			a: newIntSet(1, 2, 3, 10),
			b: newIntSet(),
			expected: map[int]bool{
				0:  false,
				1:  true,
				2:  true,
				3:  true,
				4:  false,
				9:  false,
				10: true,
				11: false,
			},
		},
		{
			a: newIntSet(),
			b: newIntSet(),
			expected: map[int]bool{
				0:     false,
				1:     false,
				2:     false,
				10:    false,
				100:   false,
				1000:  false,
				10000: false,
			},
		},
	}

	for _, test := range tests {
		test.a.UnionWith(test.b)
		for x, expected := range test.expected {
			if got := test.a.Has(x); got != expected {
				t.Errorf("Has(%d) expected: %v, but got: %v", x, expected, got)
			}
		}
	}
}

func TestIntSetString(t *testing.T) {
	tests := []struct {
		intSet   *IntSet
		expected string
	}{
		{
			intSet:   newIntSet(),
			expected: "{}",
		},
		{
			intSet:   newIntSet(1),
			expected: "{1}",
		},
		{
			intSet:   newIntSet(1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144),
			expected: "{1 2 3 5 8 13 21 34 55 89 144}",
		},
	}

	for _, test := range tests {
		if got := test.intSet.String(); got != test.expected {
			t.Errorf("intSet.String() expected: %q, but got: %q", test.expected, got)
		}
	}
}
