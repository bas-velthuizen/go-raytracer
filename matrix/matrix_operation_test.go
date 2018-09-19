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
	// Expected
	detA := 532.0
	cofA23 := -160.0
	bWanted32 := -160.0/532.0
	cofA32 := 105.0
	bWanted23 := 105.0/532.0
	bWanted := NewMatrix([][]float64{
		{  0.21805,  0.45113,  0.24060, -0.04511 },
		{ -0.80827, -1.45677, -0.44361,  0.52068 },
		{ -0.07895, -0.22368, -0.05263,  0.19737 },
		{ -0.52256, -0.81391, -0.30075,  0.30639 },
	})
	// Then
	da := a.Determinant()
	if detA != da {
		t.Errorf("det( %v ) = %9.6f , expected %9.6f", a, da, detA)
	}
	// And
	ca23 := a.Cofactor(2, 3)
	if cofA23 != ca23 {
		t.Errorf("cof( %v, 2, 3 ) = %9.6f , expected %9.6f", a, ca23, cofA23)
	}
	// And
	b32 := b.Get(3, 2)
	if bWanted32 != b32 {
		t.Errorf("b[ 3, 2] = %9.6f , expected %9.6f", b32, bWanted32)
	}
	// And
	ca32 := a.Cofactor(3, 2)
	if cofA32 != ca32 {
		t.Errorf("cof( %v, 3, 2 ) = %9.6f , expected %9.6f", a, ca32, cofA32)
	}
	// And
	b23 := b.Get(2, 3)
	if bWanted23 != b23 {
		t.Errorf("b[ 2, 3] = %9.6f , expected %9.6f", b23, bWanted23)
	}
	// And
	if !bWanted.Equals(*b) {
		t.Errorf("Inverse( %v ) = %v , expected %v", a, b, bWanted)
	}
}

// Scenario: Calculating the inverse of another matrix
// Given the following 4x4 matrix A:
// | 8 | -5 | 9 | 2 |
// | 7 | 5 | 6 | 1 |
// | -6 | 0 | 9 | 6 |
// | -3 | 0 | -9 | -4 |
// Then inverse(A) is the following 4x4 matrix:
// | -0.15385 | -0.15385 | -0.28205 | -0.53846 |
// | -0.07692 | 0.12308 | 0.02564 | 0.03077 |
// | 0.35897 | 0.35897 | 0.43590 | 0.92308 |
// | -0.69231 | -0.69231 | -0.76923 | -1.92308 |
func Test_Calculate_the_Inverse_of_Another_Matrix(t *testing.T) {
	// Given
	a := NewMatrix([][]float64{
		{  8 , -5,   9,  2 },
		{  7 ,  5,   6,  1 },
		{ -6,   0,   9,  6 },
		{ -3,   0,  -9, -4 },
	})
	// Expected
	wanted := NewMatrix([][]float64{
		{ -0.15385, -0.15385, -0.28205, -0.53846 },
		{ -0.07692,  0.12308,  0.02564,  0.03077 },
		{  0.35897,  0.35897,  0.43590,  0.92308 },
		{ -0.69231, -0.69231, -0.76923, -1.92308 },
	})
	// When
	b := a.Inverse()
	// Then
	if !wanted.Equals(*b) {
		t.Errorf("Inverse( %v ) = %v , expected %v", a, b, wanted)
	}
}
// Scenario: Calculating the inverse of a third matrix
// Given the following 4x4 matrix A:
// | 9 | 3 | 0 | 9 |
// | -5 | -2 | -6 | -3 |
// | -4 | 9 | 6 | 4 |
// | -7 | 6 | 6 | 2 |
// Then inverse(A) is the following 4x4 matrix:
// | -0.04074 | -0.07778 | 0.14444 | -0.22222 |
// | -0.07778 | 0.03333 | 0.36667 | -0.33333 |
// | -0.02901 | -0.14630 | -0.10926 | 0.12963 |
// | 0.17778 | 0.06667 | -0.26667 | 0.33333 |
func Test_Calculate_the_Inverse_of_a_Third_Matrix(t *testing.T) {
	// Given
	a := NewMatrix([][]float64{
		{  9,  3,  0,  9 },
		{ -5, -2, -6, -3 },
		{ -4,  9,  6,  4 },
		{ -7,  6,  6,  2 },
	})
	// Expected
	wanted := NewMatrix([][]float64{
		{ -0.04074, -0.07778,  0.14444, -0.22222 },
		{ -0.07778,  0.03333,  0.36667, -0.33333 },
		{ -0.02901, -0.14630, -0.10926,  0.12963 },
		{  0.17778,  0.06667, -0.26667,  0.33333 },
	})
	// When
	b := a.Inverse()
	// Then
	if !wanted.Equals(*b) {
		t.Errorf("Inverse( %v ) = %v , expected %v", a, b, wanted)
	}
}

// Scenario: Multiplying a product by its inverse
// Given the following 4x4 matrix A:
// |  3 | -9 |  7 |  3 |
// |  3 | -8 |  2 | -9 |
// | -4 |  4 |  4 |  1 |
// | -6 |  5 | -1 |  1 |
// And the following 4x4 matrix B:
// | 8 |  2 | 2 | 2 |
// | 3 | -1 | 7 | 0 |
// | 7 |  0 | 5 | 4 |
// | 6 | -2 | 0 | 5 |
// And C ← A * B
// Then C * inverse(B) = A
func Test_Multiplying_a_Product_by_its_Inverse(t *testing.T) {
	// Given
	a := NewMatrix([][]float64{
		{  3, -9,  7,  3 },
		{  3, -8,  2, -9 },
		{ -4,  4,  4,  1 },
		{ -6,  5, -1,  1 },
	})
	b := NewMatrix([][]float64{
		{ 8,  2, 2, 2 },
		{ 3, -1, 7, 0 },
		{ 7,  0, 5, 4 },
		{ 6, -2, 0, 5 },
	})
	// When
	c := a.Multiply(*b)
	// And
	d := c.Multiply(*b.Inverse())
	// Then
	if !d.Equals(*a) {
		t.Errorf("a * b * inverse(b) = %v, wanted %v", d, a)
	}

}

// Scenario: Inverting the Identity Matrix
// Given the following 4x4 matrix I:
// | 1 | 0 | 0 | 0 |
// | 0 | 1 | 0 | 0 |
// | 0 | 0 | 1 | 0 |
// | 0 | 0 | 0 | 1 |
// Then inverse(I) = I
func Test_Inverting_the_Identity_Matrix(t *testing.T) {
	// Given
	i := Identity(4)
	// When
	j := i.Inverse()
	// Then
	if !j.Equals(*i) {
		t.Errorf("i * Inverse i = %v, wanted %v", j, i)
	}
}

// Scenario: Multiplying a product by its inverse
// Given the following 4x4 matrix A:
// |  3 | -9 |  7 |  3 |
// |  3 | -8 |  2 | -9 |
// | -4 |  4 |  4 |  1 |
// | -6 |  5 | -1 |  1 |
// And the following 4x4 matrix I:
// Then A * inverse(A) = I
func Test_Multiplying_a_Matrix_by_its_Inverse(t *testing.T) {
	// Given
	a := NewMatrix([][]float64{
		{  3, -9,  7,  3 },
		{  3, -8,  2, -9 },
		{ -4,  4,  4,  1 },
		{ -6,  5, -1,  1 },
	})
	// Expected
	i := Identity(4)
	// When
	b := a.Multiply(*a.Inverse())
	// Then
	if !i.Equals(*b) {
		t.Errorf("a * inverse(a) = %v, wanted %v", b, i)
	}
}

// Scenario: Transposing the Inverse
// Given the following 4x4 matrix A:
// |  3 | -9 |  7 |  3 |
// |  3 | -8 |  2 | -9 |
// | -4 |  4 |  4 |  1 |
// | -6 |  5 | -1 |  1 |
// And B ← Inverse(Transpose(A))
// And C ← Transpose(Inverse(A))
// Then B = C
func Test_Transposing_and_Inverting_a_Matrix(t *testing.T) {
	// Given
	a := NewMatrix([][]float64{
		{  3, -9,  7,  3 },
		{  3, -8,  2, -9 },
		{ -4,  4,  4,  1 },
		{ -6,  5, -1,  1 },
	})
	// When
	b := a.Transpose().Inverse()
	// And
	c := a.Inverse().Transpose()
	// Then
	if !c.Equals(*b) {
		t.Errorf("inverse(transpose(a)) = %v , transpose(inverse(a)) = %v, should be equal", b, c)
	}
}