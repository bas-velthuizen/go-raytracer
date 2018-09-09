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
