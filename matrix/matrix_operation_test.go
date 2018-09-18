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

// Scenario: Testing an invertible matrix for invertibility
// Given the following 4x4 matrix A:
// | 6| 4| 4| 4|
// | 5| 5| 7| 6|
// | 4|-9| 3|-7|
// | 9| 1| 7|-6|
// Then determinant(A) = -2120
// And A is invertible
func Test_Invertible_Matric_For_Invertibility(t *testing.T) {
	// Given
	a := NewMatrix([][]float64{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	})
	// Expected
	wantedD := -2120.0
	// When
	d := a.Determinant()
	// Then
	if d != wantedD {
		t.Errorf("Determinant( %v ) = %9.6f, wanted %f", a, d, wantedD)
	}
	if !a.IsInvertible() {
		t.Errorf("%v is not Invertible, but should be", a)
	}
}

// Scenario: Testing a non-invertible matrix for invertibility
// Given the following 4x4 matrix A:
// |-4| 2|-2|-3|
// | 9| 6| 2| 6|
// | 0|-5| 1|-5|
// | 0| 0| 0| 0|
// Then determinant(A) = 0
// And A is not invertible
func Test_Noninvertible_Matric_For_Invertibility(t *testing.T) {
	// Given
	a := NewMatrix([][]float64{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	})
	// Expected
	wantedD := 0.0
	// When
	d := a.Determinant()
	// Then
	if d != wantedD {
		t.Errorf("Determinant( %v ) = %9.6f, wanted %f", a, d, wantedD)
	}
	if a.IsInvertible() {
		t.Errorf("%v is Invertible, but should not be", a)
	}
}

// Scenario: Calculating the inverse of a matrix
// Given the following 4x4 matrix A:
// |-5| 2| 6|-8|
// | 1|-5| 1| 8|
// | 7| 7|-6|-7|
// | 1|-3| 7| 4|
// And B ← inverse(A)
// Then determinant(A) = 532
// And cofactor(A, 2, 3) = -160
// And B[3,2] = -160/532
// And cofactor(A, 3, 2) = 105
// And B[2,3] = 105/532
// And B is the following 4x4 matrix:
//       |  0.21805 |  0.45113 |  0.24060 | -0.04511 |
//       | -0.80827 | -1.45677 | -0.44361 |  0.52068 |
//       | -0.07895 | -0.22368 | -0.05263 |  0.19737 |
//       | -0.52256 | -0.81391 | -0.30075 |  0.30639 |
func Test_Calculate_the_Inverse_of_a_Matrix(t *testing.T) {
	// Given
	a := NewMatrix([][]float64{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	})
	// When
	b := a.Inverse()

}
