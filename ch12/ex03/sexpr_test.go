package sexpr

import (
	"testing"
)

func TestBool(t *testing.T) {
	tests := []struct {
		v    bool
		want string
	}{
		{true, "t"},
		{false, "nil"},
	}
	for _, test := range tests {
		data, err := Marshal(test.v)
		if err != nil {
			t.Errorf("Marshal(%v): %s", test.v, err)
		}
		if string(data) != test.want {
			t.Errorf("Marshal(%v) got %s, wanted %s", test.v, data, test.want)
		}
	}
}

func TestFloat32(t *testing.T) {
	tests := []struct {
		v    float32
		want string
	}{
		{3.2e9, "3.2e+09"},
		{1.0, "1"},
		{0, "0"},
	}
	for _, test := range tests {
		data, err := Marshal(test.v)
		if err != nil {
			t.Errorf("Marshal(%v): %s", test.v, err)
		}
		if string(data) != test.want {
			t.Errorf("Marshal(%v) got %s, wanted %s", test.v, data, test.want)
		}
	}
}

func TestComplex64(t *testing.T) {
	tests := []struct {
		v    complex64
		want string
	}{
		{0 + 0i, "#C(0 0)"},
		{3 - 2i, "#C(3 -2)"},
		{-1e9 + -2.2e9i, "#C(-1e+09 -2.2e+09)"},
	}
	for _, test := range tests {
		data, err := Marshal(test.v)
		if err != nil {
			t.Errorf("Marshal(%v): %s", test.v, err)
		}
		if string(data) != test.want {
			t.Errorf("Marshal(%v) got %s, wanted %s", test.v, data, test.want)
		}
	}
}

func TestInterface(t *testing.T) {
	type Interface interface{}
	type Wrapper struct {
		i Interface
	}
	w := Wrapper{3}
	data, err := Marshal(w)
	if err != nil {
		t.Errorf("Marshal(%s): %s", w, err)
	}
	want := `((i ("sexpr.Interface" 3)))`
	if string(data) != want {
		t.Errorf("Marshal(%s) got %s, wanted %s", w, data, want)
	}
}
