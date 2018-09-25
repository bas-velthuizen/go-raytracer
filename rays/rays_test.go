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
