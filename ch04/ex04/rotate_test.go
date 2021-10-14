package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	s := []int{1, 2, 3}
	rotate(s)
	want := []int{2, 3, 1}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("got %v, want %v", s, want)
	}
}
