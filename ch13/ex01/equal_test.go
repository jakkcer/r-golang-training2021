package equal

import (
	"testing"
)

func TestEqual(t *testing.T) {
	for _, test := range []struct {
		x, y interface{}
		want bool
	}{
		// basic types
		{0, 0, true},
		{1000000000.9999, 1000000000.0, true},
		{1000000001, 1000000000, false},
		{uint(1000000011), uint(1000000011), true},
		{1, 1, true},
		{1, 2, false},
		{1, 1.0, false},
	} {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("Equal(%#v, %#v) = %t", test.x, test.y, !test.want)
		}
	}
}
