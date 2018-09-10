package colors

import "math"

const epsilon = 1e-5

type Color struct {
	Red float64
	Green float64
	Blue float64
}

// Equals checks if two Colors are equal
func (c Color) Equals(other Color) bool {
	return math.Abs(c.Red-other.Red) < epsilon &&
		math.Abs(c.Green-other.Green) < epsilon &&
		math.Abs(c.Blue-other.Blue) < epsilon
}

// Add adds a Color to the current Color
func (c Color) Add(other Color) Color {
	return Color{c.Red + other.Red, c.Green + other.Green, c.Blue + other.Blue}
}

// Subtract subtracts a Color from the current Color
func (c Color) Subtract(other Color) Color {
	return Color{c.Red - other.Red, c.Green - other.Green, c.Blue - other.Blue}
}

// Multiply multiplies a Color with a scalar
func (c Color) Multiply(factor float64) Color {
	return Color{c.Red * factor, c.Green * factor, c.Blue * factor}
}

// Blend multiplies a Color with another Color
func (c Color) Blend(other Color) Color {
	return Color{c.Red * other.Red, c.Green * other.Green, c.Blue * other.Blue}
}
