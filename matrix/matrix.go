package matrix

import (
	"log"
)

// Matrix is a two dimensional array
type Matrix struct {
	width  int
	height int
	data   []float64
}

// NewMatrix constructs a new Matrix from the provided data
func NewMatrix(data [][]float64) *Matrix {
	if len(data) == 0 {
		return &Matrix{width: 0, height: 0, data: []float64{}}
	}
	result := Matrix{width: len(data[0]), height: len(data), data: make([]float64, len(data)*len(data[0]))}

	for y := 0; y < result.height; y++ {
		for x := 0; x < result.width; x++ {
			result.data[y*result.width+x] = data[x][y]
			log.Printf("(%d, %d) = %f", x, y, data[x][y])
		}
	}

	return &result
}

// Get gets the value on the specified x,y position
func (m Matrix) Get(x int, y int) float64 {
	return m.data[y*m.width+x]
}
