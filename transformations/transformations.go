package transformations

import (
	"math"

	"github.com/bas-velthuizen/go-raytracer/matrix"
)

// Translation creates a new translation matrix
func Translation(x, y, z float64) *matrix.Matrix {
	t := matrix.Identity(4)
	t.Set(0, 3, x)
	t.Set(1, 3, y)
	t.Set(2, 3, z)
	return t
}

// Scaling creates a new scaling matrix
func Scaling(x, y, z float64) *matrix.Matrix {
	t := matrix.Identity(4)
	t.Set(0, 0, x)
	t.Set(1, 1, y)
	t.Set(2, 2, z)
	return t
}

// RotateX creates a new transformation matrix for rotation around the x-axis
// r is in radians
func RotationX(r float64) *matrix.Matrix {
	t := matrix.Identity(4)
	t.Set(1, 1, math.Cos(r))
	t.Set(1, 2, -math.Sin(r))
	t.Set(2, 1, math.Sin(r))
	t.Set(2, 2, math.Cos(r))
	return t
}
