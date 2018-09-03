package lib

import "math"

const tolerance = 1e-9

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
	return !t.IsPoint()
}

// Equals checks if two Tuples are equal
func (t Tuple) Equals(other Tuple) bool {
	return math.Abs(t.X-other.X) < tolerance &&
		math.Abs(t.Y-other.Y) < tolerance &&
		math.Abs(t.Z-other.Z) < tolerance &&
		math.Abs(t.W-other.W) < tolerance
}

// Point creates a new Point type Tuple
func Point(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

// Vector creates a new Vector type Tuple
func Vector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}
