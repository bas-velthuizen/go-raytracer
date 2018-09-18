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

// Scenario: Calculating a cofactor of a 3x3 matrix
// Given the following 3x3 matrix A:
// | 3 |  5 |  0 |
// | 2 | -1 | -7 |
// | 6 | -1 |  5 |
// Then minor(A, 0, 0) = -12
// And cofactor(A, 0, 0) = -12
// And minor(A, 1, 0) = 25
// And cofactor(A, 1, 0) = -25
func Test_Calculate_a_Cofactor_of_a_3x3_Matrix(t *testing.T) {
	// When
	a := NewMatrix([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	// Then
	ma := a.Minor(0, 0)
	if ma != -12.0 {
		t.Errorf("Minor( %v, 0, 0 ) = %f, wanted %f", a, ma, -12.0)
	}
	// And
	ca := a.Cofactor(0, 0)
	if ca != -12.0 {
		t.Errorf("Cofactor( %v, 0, 0 ) = %f, wanted %f", a, ca, -12.0)
	}
	// And
	mb := a.Minor(1, 0)
	if mb != 25.0 {
		t.Errorf("Minor( %v, 1, 0 ) = %f, wanted %f", a, mb, 25.0)
	}
	// And
	cb := a.Cofactor(1, 0)
	if cb != -25.0 {
		t.Errorf("Cofactor( %v, 1, 0 ) = %f, wanted %f", a, cb, -25.0)
	}
}

// Scenario: Calculating the determinant of a 3x3 matrix
// Given the following 3x3 matrix A:
// |  1 | 2 |  6 |
// | -5 | 8 | -4 |
// |  2 | 6 |  4 |
// Then cofactor(A, 0, 0) = 56
// And cofactor(A, 0, 1) = 12
// And cofactor(A, 0, 2) = -46
// And determinant(A) = -196
func Test_Calculate_Determinant_of_3x3_Matrix(t *testing.T) {
	// Given
	a := NewMatrix([][]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	})
	// Expected
	wantedCf1 := 56.0
	wantedCf2 := 12.0
	wantedCf3 := -46.0
	wantedD := -196.0
	// Then
	cf1 := a.Cofactor(0, 0)
	if wantedCf1 != cf1 {
		t.Errorf("Cofactor( %v, 0, 0 ) = %9.6f, wanted %9.6f", a, cf1, wantedCf1)
	}
	// And
	cf2 := a.Cofactor(0, 1)
	if wantedCf2 != cf2 {
		t.Errorf("Cofactor( %v, 0, 1 ) = %9.6f, wanted %9.6f", a, cf2, wantedCf2)
	}
	// And
	cf3 := a.Cofactor(0, 2)
	if wantedCf3 != cf3 {
		t.Errorf("Cofactor( %v, 0, 2 ) = %9.6f, wanted %9.6f", a, cf3, wantedCf3)
	}
	// And
	d := a.Determinant()
	if wantedD != d {
		t.Errorf("Determinant( %v ) = %9.6f, wanted %9.6f", a, d, wantedD)
	}
}

// Scenario: Calculating the determinant of a 4x4 matrix
// Given the following 4x4 matrix A:
// | -2 | -8 |  3 |  5 |
// | -3 |  1 |  7 |  3 |
// |  1 |  2 | -9 |  6 |
// | -6 |  7 |  7 | -9 |
// Then cofactor(A, 0, 0) = 690
// And cofactor(A, 0, 1) = 447
// And cofactor(A, 0, 2) = 210
// And cofactor(A, 0, 3) = 51
// And determinant(A) = -4071
func Test_Calculate_Determinant_of_4x4_Matrix(t *testing.T) {
	// Given
	a := NewMatrix([][]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	})
	// Expected
	wantedCf1 := 690.0
	wantedCf2 := 447.0
	wantedCf3 := 210.0
	wantedCf4 := 51.0
	wantedD := -4071.0
	// Then
	cf1 := a.Cofactor(0, 0)
	if wantedCf1 != cf1 {
		t.Errorf("Cofactor( %v, 0, 0 ) = %9.6f, wanted %9.6f", a, cf1, wantedCf1)
	}
	// And
	cf2 := a.Cofactor(0, 1)
	if wantedCf2 != cf2 {
		t.Errorf("Cofactor( %v, 0, 1 ) = %9.6f, wanted %9.6f", a, cf2, wantedCf2)
	}
	// And
	cf3 := a.Cofactor(0, 2)
	if wantedCf3 != cf3 {
		t.Errorf("Cofactor( %v, 0, 2 ) = %9.6f, wanted %9.6f", a, cf3, wantedCf3)
	}
	// And
	cf4 := a.Cofactor(0, 3)
	if wantedCf4 != cf4 {
		t.Errorf("Cofactor( %v, 0, 3 ) = %9.6f, wanted %9.6f", a, cf4, wantedCf4)
	}
	// And
	d := a.Determinant()
	if wantedD != d {
		t.Errorf("Determinant( %v ) = %9.6f, wanted %9.6f", a, d, wantedD)
	}
}
