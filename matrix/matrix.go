package matrix

import (
	"log"
)

// Matrix is a two dimensional array
type Matrix struct {
	size  int
	data   []float64
}

// NewMatrix constructs a new square Matrix from the provided data
func NewMatrix(data [][]float64) *Matrix {
	size := len(data)
	if size == 0 {
		return &Matrix{size: 0, data: []float64{}}
	}
	result := Matrix{size: size, data: make([]float64, size * size)}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			result.data[y*size+x] = data[x][y]
			log.Printf("(%d, %d) = %f", x, y, data[x][y])
		}
	}

	return &result
}

// Get gets the value on the specified x,y position
func (m Matrix) Get(x int, y int) float64 {
	return m.data[y*m.size+x]
}

// Equals checks if the Matrix is equal to another Matrix
func (m Matrix)Equals(other Matrix) bool {
	if m.size != other.size {
		return false
	}
	for row := 0; row < m.size; row++ {
		for col := 0; col < m.size; col++ {
			if m.data[row*m.size+col] != other.data[row*m.size+col] {
				return false
			}
		}
	}
	return true
}
