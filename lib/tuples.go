package lib

// Tuple models a point (w = 1.0) or vector (w = 0.0)
type Tuple struct {
	X float32
	Y float32
	Z float32
	W float32
}

// IsPoint checks if the Tuple is a Point
func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

// IsVector checks if the Tuple is a Vector
func (t Tuple) IsVector() bool {
	return !t.IsPoint()
}
