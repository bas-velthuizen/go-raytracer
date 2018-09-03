package lib

import (
	"reflect"
	"testing"
)

func TestConcat(t *testing.T) {
	// Given
	a := []int{1, 2, 3}
	b := []int{3, 4, 5}

	// When
	c := Concat(a, b)
	wanted := []int{1, 2, 3, 3, 4, 5}

	// Then
	if !reflect.DeepEqual(wanted, c) {
		t.Errorf("Concat(%q, %q) == %q, want %q", a, b, c, wanted)
	}
}
