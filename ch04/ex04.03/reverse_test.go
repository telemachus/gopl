package ch04_test

import (
	"ch04"
	"testing"
)

func TestReverse(t *testing.T) {
	actual := [3]int{1,2,3}
	expected := [3]int{3, 2, 1}
	ch04.Reverse(&actual)
	if expected != actual {
		t.Errorf("expected %v; actual %v", expected, actual)
	}
}
