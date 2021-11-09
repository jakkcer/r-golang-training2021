package number

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{-1}, -1},
		{[]int{1, 2}, 2},
		{[]int{-1, -2}, -1},
		{[]int{11, -22, 33, -44, 55}, 55},
		{[]int{-100, 0, 100}, 100},
	}

	for _, test := range tests {
		got := max(test.in...)
		if got != test.want {
			t.Errorf("max(%v) = %d, but want %d", test.in, got, test.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{-1}, -1},
		{[]int{1, 2}, 1},
		{[]int{-1, -2}, -2},
		{[]int{11, -22, 33, -44, 55}, -44},
		{[]int{-100, 0, 100}, -100},
	}

	for _, test := range tests {
		got := min(test.in...)
		if got != test.want {
			t.Errorf("min(%v) = %d, but want %d", test.in, got, test.want)
		}
	}
}

func TestArgsRequiredMax(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{}, 0},
		{[]int{-100, 0, 100}, 100},
	}

	for _, test := range tests {
		gotMax, gotError := argsRequiredMax(test.in...)
		if gotError == nil {
			if gotMax != test.want {
				t.Errorf("argsRequiredMax(%v) = %d, but want %d", test.in, gotMax, test.want)
			}
		} else {
			if len(test.in) != 0 {
				t.Errorf("argsRequiredMax(%v) should not return error, but returned %v", test.in, gotError)
			}
		}
	}
}

func TestArgsRequiredMin(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{}, 0},
		{[]int{-100, 0, 100}, -100},
	}

	for _, test := range tests {
		gotMin, gotError := argsRequiredMin(test.in...)
		if gotError == nil {
			if gotMin != test.want {
				t.Errorf("argsRequiredMin(%v) = %d, but want %d", test.in, gotMin, test.want)
			}
		} else {
			if len(test.in) != 0 {
				t.Errorf("argsRequiredMin(%v) should not return error, but returned %v", test.in, gotError)
			}
		}
	}
}
