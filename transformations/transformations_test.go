package transformations

import (
	"math"
	"testing"

	"github.com/bas-velthuizen/go-raytracer/tuples"
)

//
// Translation
//

// Scenario: Multiplying by a translation matrix
// Given transform ← translation(5, -3, 2)
// And p ← point(-3, 4, 5)
// Then transform * p = point(2, 1, 7)
func Test_Multiplying_By_a_Translation_Matrix(t *testing.T) {
	// Given
	transform := Translation(5, -3, 2)
	// And
	p := tuples.Point(-3, 4, 5)
	// Expected
	wanted := tuples.Point(2, 1, 7)
	// When
	r := transform.MultiplyTuple(p)
	// Then
	if !wanted.Equals(*r) {
		t.Errorf("%v * %v = %v, expected %v", transform, p, r, wanted)
	}
}

// Scenario: Multiplying by the inverse of a translation matrix
// Given transform ← translation(5, -3, 2)
// And inv ← inverse(transform)
// And p ← point(-3, 4, 5)
// Then inv * p = point(-8, 7, 3)
func Test_Multiplying_By_the_Inverse_Of_a_Translation_Matrix(t *testing.T) {
	// Given
	transform := Translation(5, -3, 2)
	// And
	inv := transform.Inverse()
	// And
	p := tuples.Point(-3, 4, 5)
	// Expected
	wanted := tuples.Point(-8, 7, 3)
	// When
	r := inv.MultiplyTuple(p)
	// Then
	if !wanted.Equals(*r) {
		t.Errorf("%v * %v = %v, expected %v", inv, p, r, wanted)
	}
}

// Scenario: Translation does not affect vectors
// Given transform ← translation(5, -3, 2)
// And v ← vector(-3, 4, 5)
// Then transform * v = v
func Test_Translation_Does_Not_Affect_Vectors(t *testing.T) {
	// Given
	transform := Translation(5, -3, 2)
	// And
	v := tuples.Vector(-3, 4, 5)
	// When
	r := transform.MultiplyTuple(v)
	// Then
	if !v.Equals(*r) {
		t.Errorf("%v * %v = %v, expected %v", transform, v, r, v)
	}
}

//
// Scaling
//

// Scenario: A scaling matrix applied to a point
// Given transform ← scaling(2, 3, 4)
// And p ← point(-4, 6, 8)
// Then transform * p = point(-8, 18, 32)
func Test_Multiplying_By_a_Scaling_Matrix(t *testing.T) {
	// Given
	transform := Scaling(2, 3, 4)
	// And
	p := tuples.Point(-4, 6, 8)
	// Expected
	wanted := tuples.Point(-8, 18, 32)
	// When
	r := transform.MultiplyTuple(p)
	// Then
	if !wanted.Equals(*r) {
		t.Errorf("%v * %v = %v, expected %v", transform, p, r, wanted)
	}
}

// Scenario: A scaling matrix applied to a vector
// Given transform ← scaling(2, 3, 4)
// And v ← vector(-4, 6, 8)
// Then transform * v = vector(-8, 18, 32)
func Test_Scaling_Does_Affect_Vectors(t *testing.T) {
	// Given
	transform := Scaling(2, 3, 4)
	// And
	v := tuples.Vector(-4, 6, 8)
	// Expected
	wanted := tuples.Vector(-8, 18, 32)
	// When
	r := transform.MultiplyTuple(v)
	// Then
	if !wanted.Equals(*r) {
		t.Errorf("%v * %v = %v, expected %v", transform, v, r, wanted)
	}
}

// Scenario: Multiplying by the inverse of a scaling matrix
// Given transform ← scaling(2, 3, 4)
// And inv ← inverse(transform)
// And v ← vector(-4, 6, 8)
// Then inv * v = vector(-2, 2, 2)
func Test_Multiplying_By_the_Inverse_Of_a_Scaling_Matrix(t *testing.T) {
	// Given
	transform := Scaling(2, 3, 4)
	// And
	inv := transform.Inverse()
	// And
	v := tuples.Vector(-4, 6, 8)
	// Expected
	wanted := tuples.Vector(-2, 2, 2)
	// When
	r := inv.MultiplyTuple(v)
	// Then
	if !wanted.Equals(*r) {
		t.Errorf("%v * %v = %v, expected %v", inv, v, r, wanted)
	}
}

// Scenario: Reflection is scaling by a negative value
// Given transform ← scaling(-1, 1, 1)
// And p ← point(2, 3, 4)
// Then transform * p = point(-2, 3, 4)
func Test_Reflection_is_Scaling_by_a_Negative_Value(t *testing.T) {
	// Given
	transform := Scaling(-1, 1, 1)
	// And
	p := tuples.Point(2, 3, 4)
	// Expected
	wanted := tuples.Point(-2, 3, 4)
	// When
	r := transform.MultiplyTuple(p)
	// Then
	if !wanted.Equals(*r) {
		t.Errorf("%v * %v = %v, expected %v", transform, p, r, wanted)
	}
}

//
// Rotation
//

// Scenario: Rotating a point around the x axis
// Given p ← point(0, 1, 0)
// And half_quarter ← rotation_x(π / 4)
// And full_quarter ← rotation_x(π / 2)
// Then half_quarter * p = point(0, √2/2, √2/2)
// And full_quarter * p = point(0, 0, 1)
func Test_Rotating_a_Point_Around_the_X_Axis(t *testing.T) {
	// Given
	p := tuples.Point(0, 1, 0)
	// And
	halfQuarter := RotationX(math.Pi / 4)
	// And
	fullQuarter := RotationX(math.Pi / 2)
	// Expected
	wantedHQ := tuples.Point(0, math.Sqrt2/2, math.Sqrt2/2)
	wantedFQ := tuples.Point(0, 0, 1)
	// When
	hq := halfQuarter.MultiplyTuple(p)
	fq := fullQuarter.MultiplyTuple(p)
	// Then
	if !wantedHQ.Equals(*hq) {
		t.Errorf("%v * %v = %v, expected %v", halfQuarter, p, hq, wantedHQ)
	}
	// And
	if !wantedFQ.Equals(*fq) {
		t.Errorf("%v * %v = %v, expected %v", fullQuarter, p, fq, wantedFQ)
	}
}

// Scenario: The inverse of an x-rotation rotates in the opposite direction
// Given v ← point(0, 1, 0)
// And half_quarter ← rotation_x(π / 4)
// And inv ← inverse(half_quarter)
// Then inv * v = point(0, √2/2, -√2/2)
func Test_The_Inverse_of_an_X_Rotation_Rotates_in_the_Opposite_Direction(t *testing.T) {
	// Given
	p := tuples.Point(0, 1, 0)
	// And
	halfQuarter := RotationX(math.Pi / 4)
	// And
	inv := halfQuarter.Inverse()
	// Expected
	wanted := tuples.Point(0, math.Sqrt2/2, -math.Sqrt2/2)
	// When
	r := inv.MultiplyTuple(p)
	// Then
	if !wanted.Equals(*r) {
		t.Errorf("inv(%v) * %v = %v, expected %v", halfQuarter, p, r, wanted)
	}
}
