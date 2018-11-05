package transformations

import (
	"math"

	"github.com/bas-velthuizen/go-raytracer/matrix"
	"github.com/bas-velthuizen/go-raytracer/tuples"
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

// RotationX creates a new transformation matrix for rotation around the x-axis
// r is in radians
func RotationX(r float64) *matrix.Matrix {
	t := matrix.Identity(4)
	t.Set(1, 1, math.Cos(r))
	t.Set(1, 2, -math.Sin(r))
	t.Set(2, 1, math.Sin(r))
	t.Set(2, 2, math.Cos(r))
	return t
}

// RotationY creates a new transformation matrix for rotation around the y-axis
// r is in radians
func RotationY(r float64) *matrix.Matrix {
	t := matrix.Identity(4)
	t.Set(0, 0, math.Cos(r))
	t.Set(0, 2, math.Sin(r))
	t.Set(2, 0, -math.Sin(r))
	t.Set(2, 2, math.Cos(r))
	return t
}

// RotationZ creates a new transformation matrix for rotation around the z-axis
// r is in radians
func RotationZ(r float64) *matrix.Matrix {
	t := matrix.Identity(4)
	t.Set(0, 0, math.Cos(r))
	t.Set(0, 1, -math.Sin(r))
	t.Set(1, 0, math.Sin(r))
	t.Set(1, 1, math.Cos(r))
	return t
}

// Shearing creates a new Transformation matrix form shearing
func Shearing(xy, xz, yx, yz, zx, zy float64) *matrix.Matrix {
	t := matrix.Identity(4)
	t.Set(0, 1, xy)
	t.Set(0, 2, xz)
	t.Set(1, 0, yx)
	t.Set(1, 2, yz)
	t.Set(2, 0, zx)
	t.Set(2, 1, zy)
	return t
}

// NewViewTransform creates a view matrix defined by the fromPoint, toPoint and upVector
func NewViewTransform(fromPoint, toPoint, upVector tuples.Tuple) matrix.Matrix {
	return *matrix.Identity(4)
}
