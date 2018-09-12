package matrix

import (
	"testing"
)

// Scenario: Constructing and inspecting a 4x4 matrix
// Given the following 4x4 matrix M:
// |  1   |  2   |  3   |  4   |
// | 5.5  | 6.5  | 7.5  | 8.5  |
// |  9   | 10   | 11   | 12   |
// | 13.5 | 14.5 | 15.5 | 16.5 |
// Then M[0,0] = 1
// And M[0,3] = 4
// And M[1,0] = 5.5
// And M[1,2] = 7.5
// And M[2,2] = 11
// And M[3,0] = 13.5
// And M[3,2] = 15.5
func Test_Construct_and_Inspect_4x4_Matrix(t *testing.T) {
	// Given
	m := NewMatrix([][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	})
	// Expected
	cases := [][]float64{
		{0, 0, 1.0},
		{0, 3, 4.0},
		{1, 0, 5.5},
		{1, 2, 7.5},
		{2, 2, 11},
		{3, 0, 13.5},
		{3, 2, 15.5},
	}
	// Then
	if m.size != 4 {
		t.Errorf("Size of Matrix = %d, wanted %d", m.size, 4)
	}
	for i := 0; i < len(cases); i++ {
		c := cases[i]
		v := m.Get(int(c[0]), int(c[1]))
		if v != c[2] {
			t.Errorf("Matrix[%d, %d] = %f, wanted %f", int(c[0]), int(c[1]), v, c[2])
		}
	}
}

// Scenario: A 2x2 matrix ought to be representable
// Given the following 2x2 matrix M:
// | -3 | 5 |
// | 1 | -2 |
// Then the size of M is 2
func Test_2x2_Matrix(t *testing.T) {
	// Given
	m := NewMatrix([][]float64{
		{-3, 5},
		{1, -2},
	})
	// Expected
	cases := [][]float64{
		{0, 0, -3},
		{0, 1, 5},
		{1, 0, 1},
		{1, 1, -2},
	}
	// Then
	if m.size != 2 {
		t.Errorf("Size of Matrix = %d, wanted %d", m.size, 2)
	}
	for i := 0; i < len(cases); i++ {
		c := cases[i]
		v := m.Get(int(c[0]), int(c[1]))
		if v != c[2] {
			t.Errorf("Matrix[%d, %d] = %f, wanted %f", int(c[0]), int(c[1]), v, c[2])
		}
	}
}

// Scenario: A 3x3 matrix ought to be representable
// Given the following 3x3 matrix M:
// | -3 | 5 | 0 |
// | 1 | -2 | -7 |
// | 0 | 1 | 1 |
// Then the size of M is 3
func Test_3x3_Matrix(t *testing.T) {
	// Given
	m := NewMatrix([][]float64{
		{-3, 5, 0},
		{1, -2, -7},
		{0, 1, 1},
	})
	// Expected
	cases := [][]float64{
		{0, 0, -3},
		{0, 1, 5},
		{1, 0, 1},
		{1, 2, -7},
		{2, 1, 1},
		{2, 2, 1},
	}
	// Then
	if m.size != 3 {
		t.Errorf("Size of Matrix = %d, wanted %d", m.size, 3)
	}
	for i := 0; i < len(cases); i++ {
		c := cases[i]
		v := m.Get(int(c[0]), int(c[1]))
		if v != c[2] {
			t.Errorf("Matrix[%d, %d] = %f, wanted %f", int(c[0]), int(c[1]), v, c[2])
		}
	}
}

// Scenario: Multiplying two matrices
// Given the following matrix A:
// | 1 | 2 | 3 | 4 |
// | 2 | 3 | 4 | 5 |
// | 3 | 4 | 5 | 6 |
// | 4 | 5 | 6 | 7 |
// And the following matrix B:
// | 0 | 1 | 2 | 4 |
// | 1 | 2 | 4 | 8 |
// | 2 | 4 | 8 | 16 |
// | 4 | 8 | 16 | 32 |
// Then A * B is the following 4x4 matrix:
// | 24 | 49 | 98 | 196 |
// | 31 | 64 | 128 | 256 |
// | 38 | 79 | 158 | 316 |
// | 45 | 94 | 188 | 376 |
func Test_Matrix_Multiplication(t *testing.T) {
	// Given
	ma := NewMatrix([][]float64{
		{ 1, 2, 3, 4 },
		{ 2, 3, 4, 5 },
		{ 3, 4, 5, 6 },
		{ 4, 5, 6, 7 },
	})
	mb := NewMatrix([][]float64{
		{ 0, 1, 2, 4 },
		{ 1, 2, 4, 8 },
		{ 2, 4, 8, 16 },
		{ 4, 8, 16, 32 },
	})
	// Expected
	wanted := NewMatrix([][]float64{
		{ 24, 49, 98, 196 },
		{ 31, 64, 128, 256 },
		{ 38, 79, 158, 316 },
		{ 45, 94, 188, 376 },
	})
	// When
	r := ma.Multiply(mb)
	// Then
	if !wanted.Equals(r) {
		t.Errorf("%v * %v = %v, wanted %v", ma, mb, r, wanted)
	}
}