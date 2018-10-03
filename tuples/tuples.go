package tuples

import (
	"fmt"
	"math"
)

// Epsilon indicates the precision of calculations with Tuples and Matrices
const Epsilon = 1e-5

// Tuple models a point (w = 1.0) or vector (w = 0.0)
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (t Tuple) String() string {
	if t.IsPoint() {
		return fmt.Sprintf("Point( %9.5f, %9.5f, %9.5f )", t.X, t.Y, t.Z)
	}
	if t.IsVector() {
		return fmt.Sprintf("Vector( %9.5f, %9.5f, %9.5f )", t.X, t.Y, t.Z)
	}
	r := fmt.Sprintf("{ %9.5f", t.X)
	r += fmt.Sprintf(", %9.5f", t.Y)
	r += fmt.Sprintf(", %9.5f", t.Z)
	r += fmt.Sprintf(", %9.5f }", t.W)
	return r
}

// IsPoint checks if the Tuple is a Point
func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

// IsVector checks if the Tuple is a Vector
func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

// Equals checks if two Tuples are equal
func (t Tuple) Equals(other Tuple) bool {
	return math.Abs(t.X-other.X) < Epsilon &&
		math.Abs(t.Y-other.Y) < Epsilon &&
		math.Abs(t.Z-other.Z) < Epsilon &&
		math.Abs(t.W-other.W) < Epsilon
}

// Add adds a Tuple to the current Tuple
func (t Tuple) Add(other Tuple) Tuple {
	return Tuple{t.X + other.X, t.Y + other.Y, t.Z + other.Z, t.W + other.W}
}

// Subtract subtracts a Tuple from the current Tuple
func (t Tuple) Subtract(other Tuple) Tuple {
	return Tuple{t.X - other.X, t.Y - other.Y, t.Z - other.Z, t.W - other.W}
}

// Negate negates a Tuple
func (t Tuple) Negate() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

// Multiply multiplies a Tuple with a scalar
func (t Tuple) Multiply(factor float64) Tuple {
	return Tuple{t.X * factor, t.Y * factor, t.Z * factor, t.W * factor}
}

// DivideBy divides a Tuple by a scalar
func (t Tuple) DivideBy(factor float64) Tuple {
	return Tuple{t.X / factor, t.Y / factor, t.Z / factor, t.W / factor}
}

// Magnitude calculates the magnitude of a vector
func (t Tuple) Magnitude() float64 {
	// Try this out: https://en.wikipedia.org/wiki/Fast_inverse_square_root
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

// Normalize normalizes a Vector
func (t Tuple) Normalize() Tuple {
	return t.DivideBy(t.Magnitude())
}

// Dot caculates dot product with another vector
func (t Tuple) Dot(other Tuple) float64 {
	return t.X*other.X +
		t.Y*other.Y +
		t.Z*other.Z +
		t.W*other.W
}

// Dot caculates dot product with another vector
func (t Tuple) Cross(other Tuple) Tuple {
	return Vector(
		t.Y*other.Z-t.Z*other.Y,
		t.Z*other.X-t.X*other.Z,
		t.X*other.Y-t.Y*other.X)
}

// Point creates a new Point type Tuple
func Point(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

// Vector creates a new Vector type Tuple
func Vector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

// Reflect returns the reflection of an incoming vector with the current vector as normal
func (t Tuple) Reflect(in Tuple) Tuple {
	return in.Subtract(t.Multiply(2 * in.Dot(t)))
}
