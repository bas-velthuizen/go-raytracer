package rays

import (
	"testing"

	"github.com/bas-velthuizen/go-raytracer/spheres"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// Scenario: An intersection encapsulates `t` and `object`
// Given s ← sphere()
// When i ← intersection(3.5, s)
// Then i.t = 3.5
// And i.object = s
func Test_an_Intersection_Encapsulates_t_and_Object(t *testing.T) {
	// Given
	s := spheres.NewUnitSphere()
	time := 3.5
	// When
	i := NewIntersection(time, s)
	// Then
	if 3.5 != i.Time {
		t.Errorf("(%v).t = %9.6f, Expected %9.6f", i, i.Time, time)
	}
	// And
	if s != i.Object {
		t.Errorf("(%v).object = %v, Expected %v", i, i.Object, s)
	}
}

// Scenario: Aggregating intersections
// Given s ← sphere()
// And i1 ← intersection(1, s)
// And i2 ← intersection(2, s)
// When xs ← intersections(i1, i2)
// Then xs.count = 2
// And xs[0].t = 1
// And xs[1].t = 2
func Test_Aggregating_Intersections(t *testing.T) {
	// Given
	s := spheres.NewUnitSphere()
	i1 := NewIntersection(1.0, s)
	i2 := NewIntersection(2.0, s)
	// When
	xs := NewIntersections([]*Intersection{i1, i2})
	// Then
	if 2 != len(*xs) {
		t.Errorf("len(%v) = %d, expected %d", xs, len(*xs), 2)
	}
	// And
	if 1.0 != (*xs)[0].Time {
		t.Errorf("xs[0] = %9.6f, expected %9.6f", 1.0, (*xs)[0].Time)
	}
	// And
	if 2.0 != (*xs)[1].Time {
		t.Errorf("xs[1] = %9.6f, expected %9.6f", 2.0, (*xs)[1].Time)
	}
}

// Scenario: Precomputing the state of an intersection
// Given ray ← ray(point(0, 0, -5), vector(0, 0, 1))
// And shape ← sphere()
// And hit ← intersection(4, shape)
// When prepare_hit(hit, ray)
// Then hit.point = point(0, 0, -1)
// And hit.eyev = vector(0, 0, -1)
// And hit.normalv = vector(0, 0, -1)
func Test_Precomputing_the_State_of_an_Intersection(t *testing.T) {
	// Given
	ray := NewRay(tuples.Point(0, 0, -5), tuples.Vector(0, 0, 1))
	// And
	shape := spheres.NewUnitSphere()
	// And
	hit := NewIntersection(4, shape)
	// When
	hit.Prepare(ray)
	// Expected
	wantedP := tuples.Point(0, 0, -1)
	wantedE := tuples.Vector(0, 0, -1)
	wantedN := tuples.Vector(0, 0, -1)
	// Then
	if !hit.Point.Equals(wantedP) {

	}
	// And
	if !hit.Eyev.Equals(wantedE) {

	}
	// And
	if !hit.Normalv.Equals(wantedN) {

	}
}
