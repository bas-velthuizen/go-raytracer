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
	for i := 0; i < len(cases); i++ {
		c := cases[i]
		v := m.Get(int(c[0]), int(c[1]))
		if v != c[2] {
			t.Errorf("Matrix[%d, %d] = %f, wanted %f", int(c[0]), int(c[1]), v, c[2])
		}
	}
}
