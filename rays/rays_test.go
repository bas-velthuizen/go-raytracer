package rays

import (
	"testing"

	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// Scenario: Creating and querying a ray
// Given origin ← point(1, 2, 3)
// And direction ← vector(4, 5, 6)
// When r ← ray(origin, direction)
// Then r.origin = origin
// And r.direction = direction
func Test_Creating_and_Querying_a_Ray(t *testing.T) {
	// Given
	origin := tuples.Point(1, 2, 3)
	// And
	direction := tuples.Vector(4, 5, 6)
	// When
	r := NewRay(origin, direction)
	// Then
	if !origin.Equals(r.origin) {
		t.Errorf("Origin of %v is %v, wanted %v", r, r.origin, origin)
	}
	// And
	if !direction.Equals(r.direction) {
		t.Errorf("direction of %v is %v, wanted %v", r, r.direction, direction)
	}
}

// Scenario: Computing a point from a distance
// Given r ← ray(point(2, 3, 4), vector(1, 0, 0))
// Then position(r, 0) = point(2, 3, 4)
// And position(r, 1) = point(3, 3, 4)
// And position(r, -1) = point(1, 3, 4)
// And position(r, 2.5) = point(4.5, 3, 4)
func Test_Computing_a_Point_from_a_Distance(t *testing.T) {
	// Given
	r := NewRay(tuples.Point(2, 3, 4), tuples.Vector(1, 0, 0))
	// Expected
	wanted0 := tuples.Point(2, 3, 4)
	wanted1 := tuples.Point(3, 3, 4)
	wantedMinus1 := tuples.Point(1, 3, 4)
	wanted2Dot5 := tuples.Point(4.5, 3, 4)
	// Then
	p0 := r.Position(0)
	if !wanted0.Equals(*p0) {
		t.Errorf("position( %v, %f ) = %v, wanted %v", r, 0.0, p0, wanted0)
	}
	// And
	p1 := r.Position(1)
	if !wanted1.Equals(*p1) {
		t.Errorf("position( %v, %f ) = %v, wanted %v", r, 1.0, p1, wanted1)
	}
	// And
	pMinus1 := r.Position(-1)
	if !wantedMinus1.Equals(*pMinus1) {
		t.Errorf("position( %v, %f ) = %v, wanted %v", r, -1.0, pMinus1, wantedMinus1)
	}
	// And
	p2Dot5 := r.Position(2.5)
	if !wanted2Dot5.Equals(*p2Dot5) {
		t.Errorf("position( %v, %f ) = %v, wanted %v", r, 2.5, p2Dot5, wanted2Dot5)
	}
}

// Scenario: A ray intersects a sphere at two points
// Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
// And s ← sphere()
// When xs ← intersect(s, r)
// Then xs.count = 2
// And xs[0] = 4
// And xs[1] = 6
func Test_a_Ray_Intersects_a_Sphere_at_Two_Points(t *testing.T) {
	// Given
	r := NewRay(tuples.Point(0, 0, -5), tuples.Vector(0, 0, 1))
	// And
	s := NewSphere(tuples.Point(0, 0, 0), 1.0)
	// Expected
	wantedCount := 2
	wanted0 := 4.0
	wanted1 := 6.0
	// When
	xs := r.Intersect(s)
	// Then
	if wantedCount != len(xs) {
		t.Errorf("intersect( %v, %v) has %d values, expected %d", s, r, len(xs), wantedCount)
	}
	// And
	if wanted0 != (*xs[0]).Time {
		t.Errorf("intersect( %v, %v)[0] is %9.6f, expected %9.6f", s, r, xs[0].Time, wanted0)
	}
	// And
	if wanted1 != (*xs[1]).Time {
		t.Errorf("intersect( %v, %v)[1] is %9.6f, expected %9.6f", s, r, xs[1].Time, wanted1)
	}
}

// Scenario: A ray intersects a sphere at a tangent
// Given r ← ray(point(0, 1, -5), vector(0, 0, 1))
// And s ← sphere()
// When xs ← intersect(s, r)
// Then xs.count = 2
// And xs[0] = 5
// And xs[1] = 5
func Test_a_Ray_Intersects_a_Sphere_at_a_Tangent(t *testing.T) {
	// Given
	r := NewRay(tuples.Point(0, 1, -5), tuples.Vector(0, 0, 1))
	// And
	s := NewSphere(tuples.Point(0, 0, 0), 1.0)
	// Expected
	wantedCount := 2
	wanted0 := 5.0
	wanted1 := 5.0
	// When
	xs := r.Intersect(s)
	// Then
	if wantedCount != len(xs) {
		t.Errorf("intersect( %v, %v) has %d values, expected %d", s, r, len(xs), wantedCount)
	}
	// And
	if wanted0 != (*xs[0]).Time {
		t.Errorf("intersect( %v, %v)[0] is %9.6f, expected %9.6f", s, r, xs[0], wanted0)
	}
	// And
	if wanted1 != (*xs[1]).Time {
		t.Errorf("intersect( %v, %v)[1] is %9.6f, expected %9.6f", s, r, xs[1], wanted1)
	}
}

// Scenario: A ray misses a sphere
// Given r ← ray(point(0, 2, -5), vector(0, 0, 1))
// And s ← sphere()
// When xs ← intersect(s, r)
// Then xs.count = 0
func Test_a_Ray_Misses_a_Sphere(t *testing.T) {
	// Given
	r := NewRay(tuples.Point(0, 2, -5), tuples.Vector(0, 0, 1))
	// And
	s := NewSphere(tuples.Point(0, 0, 0), 1.0)
	// Expected
	wantedCount := 0
	// When
	xs := r.Intersect(s)
	// Then
	if wantedCount != len(xs) {
		t.Errorf("intersect( %v, %v) has %d values, expected %d", s, r, len(xs), wantedCount)
	}
}

// Scenario: A ray originates inside a sphere
// Given r ← ray(point(0, 0, 0), vector(0, 0, 1))
// And s ← sphere()
// When xs ← intersect(s, r)
// Then xs.count = 2
// And xs[0] = -1
// And xs[1] = 1
func Test_a_Ray_Originates_Inside_a_Sphere(t *testing.T) {
	// Given
	r := NewRay(tuples.Point(0, 0, 0), tuples.Vector(0, 0, 1))
	// And
	s := NewSphere(tuples.Point(0, 0, 0), 1.0)
	// Expected
	wantedCount := 2
	wanted0 := -1.0
	wanted1 := 1.0
	// When
	xs := r.Intersect(s)
	// Then
	if wantedCount != len(xs) {
		t.Errorf("intersect( %v, %v) has %d values, expected %d", s, r, len(xs), wantedCount)
	}
	// And
	if wanted0 != (*xs[0]).Time {
		t.Errorf("intersect( %v, %v)[0] is %9.6f, expected %9.6f", s, r, xs[0], wanted0)
	}
	// And
	if wanted1 != (*xs[1]).Time {
		t.Errorf("intersect( %v, %v)[1] is %9.6f, expected %9.6f", s, r, xs[1], wanted1)
	}
}

// Scenario: A sphere is behind a ray
// Given r ← ray(point(0, 0, 5), vector(0, 0, 1))
// And s ← sphere()
// When xs ← intersect(s, r)
// Then xs.count = 2
// And xs[0] = -6
// And xs[1] = -4
func Test_a_Sphere_is_Behind_a_Ray(t *testing.T) {
	// Given
	r := NewRay(tuples.Point(0, 0, 5), tuples.Vector(0, 0, 1))
	// And
	s := NewSphere(tuples.Point(0, 0, 0), 1.0)
	// Expected
	wantedCount := 2
	wanted0 := -6.0
	wanted1 := -4.0
	// When
	xs := r.Intersect(s)
	// Then
	if wantedCount != len(xs) {
		t.Errorf("intersect( %v, %v) has %d values, expected %d", s, r, len(xs), wantedCount)
	}
	// And
	if wanted0 != (*xs[0]).Time {
		t.Errorf("intersect( %v, %v)[0] is %9.6f, expected %9.6f", s, r, xs[0], wanted0)
	}
	// And
	if wanted1 != (*xs[1]).Time {
		t.Errorf("intersect( %v, %v)[1] is %9.6f, expected %9.6f", s, r, xs[1], wanted1)
	}
}

// Scenario: Intersect sets the object on the intersection
// Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
// And s ← sphere()
// When xs ← intersect(s, r)
// Then xs.count = 2
// And xs[0].object = s
// And xs[1].object = s
func Test_Intersects_Sets_the_Object_on_the_Intersection(t *testing.T) {
	// Given
	r := NewRay(tuples.Point(0, 0, -5), tuples.Vector(0, 0, 1))
	// And
	s := NewSphere(tuples.Point(0, 0, 0), 1.0)
	// Expected
	wantedCount := 2
	// When
	xs := r.Intersect(s)
	// Then
	if wantedCount != len(xs) {
		t.Errorf("intersect( %v, %v) has %d values, expected %d", s, r, len(xs), wantedCount)
	}
	// And
	if s != (*xs[0]).Object {
		t.Errorf("intersect( %v, %v)[0] strikes %p, expected %p", s, r, xs[0].Object, s)
	}
	// And
	if s != (*xs[1]).Object {
		t.Errorf("intersect( %v, %v)[1] strikes %p, expected %p", s, r, xs[1].Object, s)
	}
}
