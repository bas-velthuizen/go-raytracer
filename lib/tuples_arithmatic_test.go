package lib

import "testing"

//Scenario: Adding two vectors
//Given a1 ← tuple(3, -2, 5, 0)
//And a2 ← tuple(-2, 3, 1, 0)
//Then a1 + a2 = tuple(1, 1, 6, 0)
func Test_Adding_two_vectors(t *testing.T) {
	// Given
	a1 := Tuple{3, -2, 5, 0}
	a2 := Tuple{-2, 3, 1, 0}
	// Expected
	wanted := Tuple{1, 1, 6, 0}
	// Then
	p := a1.Add(a2)
	if !p.Equals(wanted) {
		t.Errorf("%v + %v = %v, want %v", a1, a2, p, wanted)
	}
	if !p.IsVector() {
		t.Errorf("%v + %v returns Vector? %t, wanted true", a1, a2, p.IsVector())
	}
}

//Scenario: Adding two points
//Given a1 ← tuple(3, -2, 5, 1)
//And a2 ← tuple(-2, 3, 1, 1)
//Then a1 + a2 = tuple(1, 1, 6, 2) ; neither a Point nor a Vector
func Test_Adding_two_points(t *testing.T) {
	// Given
	a1 := Tuple{3, -2, 5, 1}
	a2 := Tuple{-2, 3, 1, 1}
	// Expected
	wanted := Tuple{1, 1, 6, 2}
	// Then
	p := a1.Add(a2)
	if !p.Equals(wanted) {
		t.Errorf("%v + %v = %v, want %v", a1, a2, p, wanted)
	}
	if p.IsPoint() {
		t.Errorf("%v + %v returns Point? %t, wanted false", a1, a2, p.IsPoint())
	}
	if p.IsVector() {
		t.Errorf("%v + %v returns Vector? %t, wanted false", a1, a2, p.IsVector())
	}
}

// Scenario: Subtracting two points
// Given p1 ← point(3, 2, 1)
// And p2 ← point(5, 6, 7)
// Then p1 - p2 = vector(-2, -4, -6)
func Test_Subtracting_two_points(t *testing.T) {
	// Given
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)
	// Expected
	wanted := Vector(-2, -4, -6)
	// Then
	p := p1.Subtract(p2)
	if !p.Equals(wanted) {
		t.Errorf("%v - %v = %v, want %v", p1, p2, p, wanted)
	}
	if p.IsPoint() {
		t.Errorf("%v + %v returns Point? %t, wanted false", p1, p2, p.IsPoint())
	}
	if !p.IsVector() {
		t.Errorf("%v + %v returns Vector? %t, wanted true", p1, p2, p.IsVector())
	}
}

// Scenario: Subtracting a vector from a point
// Given p ← point(3, 2, 1)
// And v ← vector(5, 6, 7)
// Then p - v = point(-2, -4, -6)
func Test_Subtracting_vector_from_point(t *testing.T) {
	// Given
	p1 := Point(3, 2, 1)
	v := Vector(5, 6, 7)
	// Expected
	wanted := Point(-2, -4, -6)
	// Then
	p := p1.Subtract(v)
	if !p.Equals(wanted) {
		t.Errorf("%v - %v = %v, want %v", p1, v, p, wanted)
	}
	if !p.IsPoint() {
		t.Errorf("%v + %v returns Point? %t, wanted true", p1, v, p.IsPoint())
	}
	if p.IsVector() {
		t.Errorf("%v + %v returns Vector? %t, wanted false", p1, v, p.IsVector())
	}
}

// Scenario: Subtracting two vectors
// Given v1 ← vector(3, 2, 1)
// And v2 ← vector(5, 6, 7)
// Then v1 - v2 = vector(-2, -4, -6)
func Test_Subtracting_two_vectors(t *testing.T) {
	// Given
	v1 := Vector(3, 2, 1)
	v2 := Vector(5, 6, 7)
	// Expected
	wanted := Vector(-2, -4, -6)
	// Then
	v := v1.Subtract(v2)
	if !v.Equals(wanted) {
		t.Errorf("%v - %v = %v, want %v", v1, v2, v, wanted)
	}
	if v.IsPoint() {
		t.Errorf("%v + %v returns Point? %t, wanted false", v1, v2, v.IsPoint())
	}
	if !v.IsVector() {
		t.Errorf("%v + %v returns Vector? %t, wanted true", v1, v2, v.IsVector())
	}
}

// Scenario: Subtracting a vector from the zero vector
// Given zero ← vector(0, 0, 0)
// And v ← vector(1, -2, 3)
// Then zero - v = vector(-1, 2, -3)
func Test_Subtracting_vector_from_zero_vector(t *testing.T) {
	// Given
	v1 := Vector(1, -2, 3)
	zero := Vector(0, 0, 0)
	// Expected
	wanted := Vector(-1, 2, -3)
	// Then
	v := zero.Subtract(v1)
	if !v.Equals(wanted) {
		t.Errorf("%v - %v = %v, want %v", zero, v1, v, wanted)
	}
	if v.IsPoint() {
		t.Errorf("%v + %v returns Point? %t, wanted false", zero, v1, v.IsPoint())
	}
	if !v.IsVector() {
		t.Errorf("%v + %v returns Vector? %t, wanted true", zero, v1, v.IsVector())
	}
}

// Scenario: Negating a tuple
// Given a ← tuple(1, -2, 3, -4)
// Then -a = tuple(-1, 2, -3, 4)
func Test_Negating_a_tuple(t *testing.T) {
	// Given
	t1 := Tuple{1, -2, 3, -4}
	// Expected
	wanted := Tuple{-1, 2, -3, 4}
	// Then
	t2 := t1.Negate()
	if !t2.Equals(wanted) {
		t.Errorf("- %v = %v, want %v", t1, t2, wanted)
	}
}

// Scenario: Multiplying a tuple by a scalar
// Given a ← tuple(1, -2, 3, -4)
// Then a * 3.5 = tuple(3.5, -7, 10.5, -14)
func Test_Multiplying_a_tuple_by_a_scalar(t *testing.T) {
	// Given
	t1 := Tuple{1, -2, 3, -4}
	// Expected
	wanted := Tuple{3.5, -7, 10.5, -14}
	// Then
	t2 := t1.Multiply(3.5)
	if !t2.Equals(wanted) {
		t.Errorf("3.5 * %v = %v, want %v", t1, t2, wanted)
	}
}

// Scenario: Multiplying a tuple by a fraction
// Given a ← tuple(1, -2, 3, -4)
// Then a * 0.5 = tuple(0.5, -1, 1.5, -2)
func Test_Multiplying_a_tuple_by_a_faraction(t *testing.T) {
	// Given
	t1 := Tuple{1, -2, 3, -4}
	// Expected
	wanted := Tuple{0.5, -1, 1.5, -2}
	// Then
	t2 := t1.Multiply(.5)
	if !t2.Equals(wanted) {
		t.Errorf(".5 * %v = %v, want %v", t1, t2, wanted)
	}
}

// Scenario: Dividing a tuple by a scalar
// Given a ← tuple(1, -2, 3, -4)
// Then a / 2 = tuple(0.5, -1, 1.5, -2)
func Test_Dividing_a_tuple_by_a_scalar(t *testing.T) {
	// Given
	t1 := Tuple{1, -2, 3, -4}
	// Expected
	wanted := Tuple{0.5, -1, 1.5, -2}
	// Then
	t2 := t1.DivideBy(2)
	if !t2.Equals(wanted) {
		t.Errorf("%v / 2 = %v, want %v", t1, t2, wanted)
	}
}
