package tuples

import "testing"

//Scenario: A tuple with w=1.0 is a point
// Given a ← tuple(4.3, -4.2, 3.1, 1.0)
// Then a.x = 4.3
// And a.y = -4.2
// And a.z = 3.1
// And a.w = 1.0
// And a is a point
// And a is not a vector
func Test_Tuple_with_w_1_is_point(t *testing.T) {
	// Given
	a := Tuple{4.3, -4.2, 3.1, 1.0}
	// Then
	if a.X != 4.3 {
		t.Errorf("a.X == %f, want %f", a.X, 4.3)
	}
	if a.Y != -4.2 {
		t.Errorf("a.Y == %f, want %f", a.Y, -4.2)
	}
	if a.Z != 3.1 {
		t.Errorf("a.Z == %f, want %f", a.Z, 3.1)
	}
	if a.W != 1.0 {
		t.Errorf("a.W == %f, want %f", a.W, 1.0)
	}
	if !a.IsPoint() {
		t.Errorf("a.IsPoint == %t, want %t", a.IsPoint(), true)
	}
	if a.IsVector() {
		t.Errorf("a.IsVector == %t, want %t", a.IsVector(), false)
	}
}

// Scenario: A tuple with w=0 is a vector
// Given a ← tuple(4.3, -4.2, 3.1, 0.0)
// Then a.x = 4.3
// And a.y = -4.2
// And a.z = 3.1
// And a.w = 0.0
// And a is not a point
// And a is a vector
func Test_Tuple_with_w_0_is_vector(t *testing.T) {
	// Given
	a := Tuple{4.3, -4.2, 3.1, 0.0}
	// Then
	if a.X != 4.3 {
		t.Errorf("a.X == %f, want %f", a.X, 4.3)
	}
	if a.Y != -4.2 {
		t.Errorf("a.Y == %f, want %f", a.Y, -4.2)
	}
	if a.Z != 3.1 {
		t.Errorf("a.Z == %f, want %f", a.Z, 3.1)
	}
	if a.W != 0.0 {
		t.Errorf("a.W == %f, want %f", a.W, 0.0)
	}
	if a.IsPoint() {
		t.Errorf("a.IsPoint == %t, want %t", a.IsPoint(), false)
	}
	if !a.IsVector() {
		t.Errorf("a.IsVector == %t, want %t", a.IsVector(), true)
	}
}

// Scenario: "point" describes tuples with w=1
// Given p ← point(4, -4, 3)
// Then p = tuple(4, -4, 3, 1)
func Test_Point_describes_tuples_with_w_1(t *testing.T) {
	// Given
	p := Point(4, -4, 3)
	// Expected
	wanted := Tuple{4, -4, 3, 1.0}
	// Then
	if !p.Equals(wanted) {
		t.Errorf("Point(4, -4, 3) == %v, want %v", p, wanted)
	}
}

// Scenario: "vector" describes tuples with w=0
// Given v ← vector(4, -4, 3)
// Then v = tuple(4, -4, 3, 0)
func Test_Vector_describes_tuples_with_w_1(t *testing.T) {
	// Given
	p := Vector(4, -4, 3)
	// Expected
	wanted := Tuple{4, -4, 3, 0.0}
	// Then
	if !p.Equals(wanted) {
		t.Errorf("Vector(4, -4, 3) == %v, want %v", p, wanted)
	}
}
