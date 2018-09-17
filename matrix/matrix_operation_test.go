package matrix

import (
	"math"
	"testing"

	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// Scenario: Calculating the determinant of a 2x2 matrix
// Given the following 2x2 matrix A:
// | 1 | 5 |
// | -3 | 2 |
// Then determinant(A) = 17
func Test_Calculating_Determinant_of_2x2_Matrix(t *testing.T) {
	// Given
	a := NewMatrix([][]float64{
		{1, 5},
		{-3, 2},
	})
	// Expected
	wanted := 17.0
	// When
	d := a.Determinant()
	// Then
	if math.Abs(wanted-d) > tuples.Epsilon {
		t.Errorf("determinant( %v )= %9.5f, wanted %9.5f", a, d, wanted)
	}
}

// Scenario: A submatrix of a 3x3 matrix is a 2x2 matrix
// Given the following 3x3 matrix A:
// |  1 | 5 |  0 |
// | -3 | 2 |  7 |
// |  0 | 6 | -3 |
// Then submatrix(A, 0, 2) is the following 2x2 matrix:
// | -3 | 2 |
// |  0 | 6 |
func Test_Submatrix_of_3x3_Matrix_is_2x2_Matrix(t *testing.T) {
	// Given
	a := NewMatrix([][]float64{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	})
	// Expected
	wanted := NewMatrix([][]float64{
		{-3, 2},
		{0, 6},
	})
	// When
	sub := a.Submatrix(0, 2)
	// Then
	if !wanted.Equals(*sub) {
		t.Errorf("submatrix(%v, 0, 2) = %v, wanted %v", a, sub, wanted)
	}
}

// Scenario: A submatrix of a 4x4 matrix is a 3x3 matrix Given the following 4x4 matrix A:
// | -6 | 1 |  1 | 6 |
// | -8 | 5 |  8 | 6 |
// | -1 | 0 |  8 | 2 |
// | -7 | 1 | -1 | 1 |
// Then submatrix(A, 2, 1) is the following 3x3 matrix:
// | -6 |  1 | 6 |
// | -8 |  8 | 6 |
// | -7 | -1 | 1 |
func Test_Submatrix_of_4x4_Matrix_is_3x3_Matrix(t *testing.T) {
	// Given
	a := NewMatrix([][]float64{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	})
	// Expected
	wanted := NewMatrix([][]float64{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	})
	// When
	sub := a.Submatrix(2, 1)
	// Then
	if !wanted.Equals(*sub) {
		t.Errorf("submatrix(%v, 0, 2) = %v, wanted %v", a, sub, wanted)
	}
}

// Scenario: Calculating a minor of a 3x3 matrix
// Given the following 3x3 matrix A:
// | 3 |  5 |  0 |
// | 2 | -1 | -7 |
// | 6 | -1 |  5 |
// And B ← submatrix(A, 1, 0)
// Then determinant(B) = 25
// And minor(A, 1, 0) = 25
func Test_Calculate_Minor_of_3x3_Matrix(t *testing.T) {
	// Given
	a := NewMatrix([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	// Expected
	wanted := 25.0
	// When
	b := a.Submatrix(1, 0)
	// Then
	bd := b.Determinant()
	if bd != wanted {
		t.Errorf("Determinant( %v ) == %f, wanted %f", b, bd, wanted)
	}
	// And
	ma := a.Minor(1, 0)
	if ma != wanted {
		t.Errorf("Minor( %v, 1, 0 ) == %f, wanted %f", a, ma, wanted)
	}
}
