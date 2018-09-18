package matrix

import (
	"fmt"
	"log"
	"math"

	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// Matrix is a two dimensional array
type Matrix struct {
	size int
	data []float64
}

// NewMatrix constructs a new square Matrix from the provided data
func NewMatrix(data [][]float64) *Matrix {
	size := len(data)
	if size == 0 {
		return &Matrix{size: 0, data: []float64{}}
	}
	result := Matrix{size: size, data: make([]float64, size*size)}

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			result.Set(row, col, data[row][col])
			log.Printf("(%d, %d) = %f", row, col, data[row][col])
		}
	}

	return &result
}

// Identity creates an identity matrix of the specified size
func Identity(size int) *Matrix {
	if size == 0 {
		return &Matrix{size: 0, data: []float64{}}
	}
	i := Matrix{size: size, data: make([]float64, size*size)}

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			v := 0.0
			if row == col {
				v = 1.0
			}
			i.Set(row, col, v)
		}
	}

	return &i
}

func (m Matrix) String() string {
	r := "{ "
	for row := 0; row < m.size; row++ {
		r += "{"
		for col := 0; col < m.size; col++ {
			if col != 0 {
				r += ", "
			}
			r += fmt.Sprintf(" %9.5f", m.Get(row, col))
		}
		r += " }"
		if row != m.size-1 {
			r += ","
		}
		r += "\n"
	}
	r += " }"
	return r
}

// Get gets the value on the specified position
func (m Matrix) Get(row int, col int) float64 {
	return m.data[row*m.size+col]
}

// Set sets the value on the specified position
func (m Matrix) Set(row int, col int, v float64) {
	m.data[row*m.size+col] = v
}

// Equals checks if the Matrix is equal to another Matrix
func (m Matrix) Equals(other Matrix) bool {
	if m.size != other.size {
		return false
	}
	for row := 0; row < m.size; row++ {
		for col := 0; col < m.size; col++ {
			if math.Abs(m.data[row*m.size+col]-other.data[row*m.size+col]) > tuples.Epsilon {
				return false
			}
		}
	}
	return true
}

// Multiply calculates the product of two matrices
func (m Matrix) Multiply(other Matrix) *Matrix {
	p := NewMatrix([][]float64{})
	p.size = m.size
	p.data = make([]float64, m.size*m.size)
	for row := 0; row < p.size; row++ {
		for col := 0; col < p.size; col++ {
			p.Set(row, col, m.rowToTuple(row).Dot(other.columnToTuple(col)))
		}
	}
	return p
}

// Multiply calculates the product of two matrices
func (m Matrix) MultiplyVector(t tuples.Tuple) *tuples.Tuple {
	p := &tuples.Tuple{}
	p.X = m.rowToTuple(0).Dot(t)
	p.Y = m.rowToTuple(1).Dot(t)
	p.Z = m.rowToTuple(2).Dot(t)
	p.W = m.rowToTuple(3).Dot(t)
	return p
}

// Transpose transposes a Matrix, ie. mirrors it along the r=c diagonal
func (m Matrix) Transpose() *Matrix {
	t := NewMatrix([][]float64{})
	t.size = m.size
	t.data = make([]float64, len(m.data))
	for row := 0; row < t.size; row++ {
		for col := 0; col < t.size; col++ {
			t.Set(row, col, m.Get(col, row))
		}
	}
	return t
}

// Determinant calculates the Determinant of a matrix
func (m Matrix) Determinant() float64 {
	if m.size == 1 {
		return m.Get(0, 0)
	}
	// if m.size == 2 {
	// 	d := m.Get(0, 0)*m.Get(1, 1) - m.Get(0, 1)*m.Get(1, 0)
	// 	return d
	// }
	fmt.Printf("m.size: %d \n", m.size)
	d := 0.0
	for col := 0; col < m.size; col++ {
		fmt.Printf("m(0, %d) = %f , cov(m, 1, %d) = %f\n", col, m.Get(0, col), col, m.Cofactor(1, col))
		d += m.Get(0, col) * m.Cofactor(0, col)
	}
	return d
}

// Submatrix returns the specified submatrix of a Matrix
// it removes the specified row an column
func (m Matrix) Submatrix(row int, col int) *Matrix {
	sub := NewMatrix([][]float64{})
	sub.size = m.size - 1
	sub.data = make([]float64, (sub.size * sub.size))
	for srow := 0; srow < m.size; srow++ {
		trow := srow
		if trow > row {
			trow--
		}
		for scol := 0; scol < m.size; scol++ {
			tcol := scol
			if tcol > col {
				tcol--
			}
			if srow != row && scol != col {
				sub.Set(trow, tcol, m.Get(srow, scol))
			}
		}
	}
	return sub
}

// Minor calculates the specified minor of a Matrix
// which means the determinant of the specified submatrix
func (m Matrix) Minor(row int, col int) float64 {
	sub := m.Submatrix(row, col)
	return sub.Determinant()
}

// Cofactor calculates the specified cofactor of a Matrix
// which means the sign-corrected determinant of the specified submatrix
func (m Matrix) Cofactor(row int, col int) float64 {
	sub := m.Submatrix(row, col)
	if (row+col)%2 == 1 {
		return -sub.Determinant()
	}
	return sub.Determinant()
}

func (m Matrix) rowToTuple(row int) tuples.Tuple {
	return tuples.Tuple{X: m.Get(row, 0), Y: m.Get(row, 1), Z: m.Get(row, 2), W: m.Get(row, 3)}
}

func (m Matrix) columnToTuple(col int) tuples.Tuple {
	return tuples.Tuple{X: m.Get(0, col), Y: m.Get(1, col), Z: m.Get(2, col), W: m.Get(3, col)}
}
