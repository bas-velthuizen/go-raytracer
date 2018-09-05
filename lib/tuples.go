package lib

import "math"

const epsilon = 1e-5

// Tuple models a point (w = 1.0) or vector (w = 0.0)
type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
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
	return math.Abs(t.X-other.X) < epsilon &&
		math.Abs(t.Y-other.Y) < epsilon &&
		math.Abs(t.Z-other.Z) < epsilon &&
		math.Abs(t.W-other.W) < epsilon
}

// Add adds a Tuple to the current Tuple
func (t Tuple) Add(other Tuple) Tuple {
	return Tuple{t.X + other.X, t.Y + other.Y, t.Z + other.Z, t.W + other.W}
}

// Point creates a new Point type Tuple
func Point(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

// Vector creates a new Vector type Tuple
func Vector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}
