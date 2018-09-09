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
