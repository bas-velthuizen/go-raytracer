package rays

import (
	"github.com/bas-velthuizen/go-raytracer/tuples"
	"testing"
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
	r := Ray(origin, direction)
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
	r := Ray(tuples.Point(2, 3, 4), tuples.Vector(1, 0, 0))
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
