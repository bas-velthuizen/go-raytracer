package matrix

import (
	"github.com/bas-velthuizen/go-raytracer/tuples"
	"math"
	"testing"
)

// Scenario: Calculating the determinant of a 2x2 matrix
// Given the following 2x2 matrix A:
// | 1 | 5 |
// | -3 | 2 |
// Then determinant(A) = 17
func Test_Calculating_Determinant_of_2x2_Matrix(t *testing.T){
	// Given
	a := NewMatrix([][]float64{
		{  1, 5 },
		{ -3, 2 },
	})
	// Expected
	wanted := 17.0
	// When
	d := a.Determinant()
	// Then
	if math.Abs(wanted - d) > tuples.Epsilon {
		t.Errorf("determinant( %v )= %9.5f, wanted %9.5f", a, d, wanted)
	}
}