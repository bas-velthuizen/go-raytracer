package colors

import "testing"

//Scenario: Colors are (red, green, blue) tuples
//Given c ← color(-0.5, 0.4, 1.7)
//Then c.red = -0.5
//And c.green = 0.4
//And c.blue = 1.7
func Test_Colors_are_RGB_Tuples(t *testing.T) {
	// Given
	c := Color{-0.5, 0.4, 1.7}
	// Then
	if c.Red != -0.5 {
		t.Errorf("c.X == %f, want %f", c.Red, -0.5)
	}
	if c.Green != 0.4 {
		t.Errorf("c.Y == %f, want %f", c.Green, 0.4)
	}
	if c.Blue != 1.7 {
		t.Errorf("c.Z == %f, want %f", c.Blue, 1.7)
	}
}

//Scenario: Adding colors
//Given c1 ← color(0.9, 0.6, 0.75)
//And c2 ← color(0.7, 0.1, 0.25)
//Then c1 + c2 = color(1.6, 0.7, 1.0)
func Test_Add_Colors(t *testing.T) {
	// Given
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	//Expected
	wanted:= Color{1.6, 0.7, 1.0}
	// Then
	r := c1.Add(c2)
	if !r.Equals(wanted) {
		t.Errorf("%v + %v = %v, want %v", c1, c2, r, wanted)
	}
}

//Scenario: Subtracting colors
//Given c1 ← color(0.9, 0.6, 0.75)
//And c2 ← color(0.7, 0.1, 0.25)
//Then c1 - c2 = color(0.2, 0.5, 0.5)
func Test_Subtract_Colors(t *testing.T) {
	// Given
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	//Expected
	wanted:= Color{0.2, 0.5, 0.5}
	// Then
	r := c1.Subtract(c2)
	if !r.Equals(wanted) {
		t.Errorf("%v - %v = %v, want %v", c1, c2, r, wanted)
	}
}

//Scenario: Multiplying a color by a scalar
//Given c ← color(0.2, 0.3, 0.4)
//Then c * 2 = color(0.4, 0.6, 0.8)
func Test_Multiplying_a_color_by_a_scalar(t *testing.T) {
	// Given
	c := Color{0.2, 0.3, 0.4}
	// Expected
	wanted := Color{0.4, 0.6, 0.8}
	// Then
	r := c.Multiply(2)
	if !r.Equals(wanted) {
		t.Errorf("2 * %v = %v, want %v", c, r, wanted)
	}
}

//Scenario: Multiplying colors (Hadamard or Schur product, blending)
//Given c1 ← color(1, 0.2, 0.4)
//And c2 ← color(0.9, 1, 0.1)
//Then c1 * c2 = color(0.9, 0.2, 0.04)
func Test_Blend_Colors(t *testing.T) {
	// Given
	c1 := Color{1, 0.2, 0.4}
	c2 := Color{0.9, 1, 0.1}
	//Expected
	wanted:= Color{0.9, 0.2, 0.04}
	// Then
	r := c1.Blend(c2)
	if !r.Equals(wanted) {
		t.Errorf("%v - %v = %v, want %v", c1, c2, r, wanted)
	}
}
