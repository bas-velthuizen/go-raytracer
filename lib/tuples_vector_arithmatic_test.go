package lib

import (
	"math"
	"testing"
)

// Scenario: Magnitude of vector(1, 0, 0)
// Given v ← vector(1, 0, 0)
// Then magnitude(v) = 1
// Scenario: Magnitude of vector(0, 1, 0)
// Given v ← vector(0, 1, 0)
// Then magnitude(v) = 1
// Scenario: Magnitude of vector(0, 0, 1)
// Given v ← vector(0, 0, 1)
// Then magnitude(v) = 1
// Scenario: Magnitude of vector(1, 2, 3)
// Given v ← vector(1, 2, 3)
// Then magnitude(v) = √14
// Scenario: Magnitude of vector(-1, -2, -3)
// Given v ← vector(-1, -2, -3)
// Then magnitude(v) = √14
func Test_Magnitude_Of_Vector(t *testing.T) {
	cases := []Tuple{Vector(1, 0, 0), Vector(0, 1, 0), Vector(0, 0, 1), Vector(1, 2, 3), Vector(-1, -2, -3)}
	wanteds := []float64{1.0, 1.0, 1.0, math.Sqrt(14), math.Sqrt(14)}
	for i := 0; i < len(cases); i++ {
		// Given
		v := cases[i]
		// Expected
		wanted := wanteds[i]
		// Then
		m := v.Magnitude()
		if wanted != m {
			t.Errorf("magnitude of %v = %v, want %v", v, m, wanted)
		}
	}
}

// Scenario: Normalizing vector(4, 0, 0) gives (1, 0, 0)
// Given v ← vector(4, 0, 0)
// Then normalize(v) = vector(1, 0, 0)
// Scenario: Normalizing vector(1, 2, 3)
// Given v ← vector(1, 2, 3)
// Then normalize(v) = approximately vector(0.26726, 0.53452, 0.80178) # vector(1/√14, 2/√14, 3/√14)
// Scenario: The magnitude of a normalized vector
// Given v ← vector(1, 2, 3)
// When norm ← normalize(v)
// Then magnitude(norm) = 1
func Test_Normalizing_a_Vector(t *testing.T) {
	cases := []Tuple{Vector(1, 0, 0), Vector(4, 0, 0), Vector(1, 2, 3)}
	wanteds := []Tuple{Vector(1, 0, 0), Vector(1, 0, 0), Vector(0.26726, 0.53452, 0.80178)}
	for i := 0; i < len(cases); i++ {
		// Given
		v := cases[i]
		// Expected
		wanted := wanteds[i]
		// Then
		m := v.Normalize()
		if !wanted.Equals(m) {
			t.Errorf("Normalize of %v = %v, want %v", v, m, wanted)
		}
		n := m.Magnitude()
		if 1.0 != n {
			t.Errorf("magnitude of %v = %v, want 1.0", m, n)
		}
	}
}

// Scenario: The dot product of two tuples
// Given a ← vector(1, 2, 3)
// And b ← vector(2, 3, 4) Then a dot b = 20
func Test_Dot_Product_of_two_Vectors(t *testing.T) {
	// Given
	t1 := Vector(1, 2, 3)
	t2 := Vector(2, 3, 4)
	// Expected
	wanted := 20.0
	// Then
	d := t1.Dot(t2)
	if d != wanted {
		t.Errorf("%v . %v = %v, want %v", t1, t2, d, wanted)
	}
}

// Scenario: Cross product of two vectors
// Given a ← vector(1, 2, 3)
// And b ← vector(2, 3, 4)
// Then a cross b = vector(-1, 2, -1)
// And b cross a = vector(1, -2, 1)
func Test_Cross_Product_of_two_Vectors(t *testing.T) {
	// Given
	t1 := Vector(1, 2, 3)
	t2 := Vector(2, 3, 4)
	// Expected
	wanted1 := Vector(-1, 2, -1)
	wanted2 := Vector(1, -2, 1)
	// Then
	v1 := t1.Cross(t2)
	v2 := t2.Cross(t1)
	if !wanted1.Equals(v1) {
		t.Errorf("%v x %v = %v, want %v", t1, t2, v1, wanted1)
	}
	if !wanted2.Equals(v2) {
		t.Errorf("%v x %v = %v, want %v", t2, t1, v2, wanted2)
	}

}
