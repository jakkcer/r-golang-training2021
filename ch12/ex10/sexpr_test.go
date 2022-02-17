package sexpr

import (
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type Interface interface{}
	type Record struct {
		B    bool
		F32  float32
		F64  float64
		C64  complex64
		C128 complex128
		I    Interface
	}
	Interfaces["sexpr.Interface"] = reflect.TypeOf(int(0))
	tests := []struct {
		s    string
		want Record
	}{
		{
			`((B t) (F32 2.5) (F64 0) (I ("sexpr.Interface" 5)))`,
			Record{true, 2.5, 0, 0, 0, Interface(5)},
		},
		{
			`((B nil) (F32 0) (F64 1.5) (I ("sexpr.Interface" 0)))`,
			Record{false, 0, 1.5, 0, 0, Interface(0)},
		},
	}
	for _, test := range tests {
		var r Record
		err := Unmarshal([]byte(test.s), &r)
		if err != nil {
			t.Errorf("Unmarshal(%q): %s", test.s, err)
		}
		if !reflect.DeepEqual(r, test.want) {
			t.Errorf("Unmarshal(%q) got %#v, wanted %#v", test.s, r, test.want)
		}
	}
}
