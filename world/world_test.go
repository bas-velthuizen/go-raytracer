package world

import (
	"testing"

	"github.com/bas-velthuizen/go-raytracer/colors"
	"github.com/bas-velthuizen/go-raytracer/spheres"
	"github.com/bas-velthuizen/go-raytracer/transformations"
	"github.com/bas-velthuizen/go-raytracer/tuples"

	"github.com/bas-velthuizen/go-raytracer/lights"
	"github.com/bas-velthuizen/go-raytracer/rays"
)

// Scenario: Creating a world
// Given w ← world()
// Then w contains no objects
// And w has no light source
func Test_Creating_a_World(t *testing.T) {
	// Given
	w := NewWorld([]spheres.Sphere{}, nil)
	// Then
	if 0 != len(w.Objects) {
		t.Errorf("%v has %d Objects, expected %d", w, len(w.Objects), 0)
	}
	// And w has no light source
	if nil != w.LightSource {
		t.Errorf("%v has LightSource %v, expected none", w, w.LightSource)
	}
}

// Scenario: The default world
// Given light ← point_light(point(-10, 10, -10), color(1, 1, 1))
// And s1 ← sphere() with:
// | color | (0.8, 1.0, 0.6) | | diffuse | 0.7 | | specular | 0.2 |
// And s2 ← sphere() with:
// | transform | scaling(0.5, 0.5, 0.5) |
// When world ← default_world()
// Then world.light = light
// And world contains s1
// And world contains s2
func Test_The_Default_World(t *testing.T) {
	// Given
	light := lights.NewPointLight(tuples.Point(-10, 10, -10), colors.White())
	// And
	s1 := spheres.NewUnitSphere()
	s1.Material.Color = colors.NewColor(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2
	// And
	s2 := spheres.NewUnitSphere()
	s2.SetTransform(transformations.Scaling(0.5, 0.5, 0.5))
	// When
	world := DefaultWorld()
	// Then
	if !light.Equals(*world.LightSource) {
		t.Errorf("%v has LightSource %v, expected %v", world, world.LightSource, light)
	}
	// And
	if !world.Contains(*s1) {
		t.Errorf("Expected %v to contain %v, but it doesn't", world, s1)
	}
	// And
	if !world.Contains(*s2) {
		t.Errorf("Expected %v to contain %v, but it doesn't", world, s2)
	}
}

// Scenario: Intersect a world with a ray
// Given world ← default_world()
// And ray ← ray(point(0, 0, -5), vector(0, 0, 1))
// When xs ← intersect_world(world, ray)
// Then xs.count = 4
// And xs[0].t = 4
// And xs[1].t = 4.5
// And xs[2].t = 5.5

func Test_Intersect_a_World_With_a_Ray(t *testing.T) {
	// Given
	world := DefaultWorld()
	// And
	ray := rays.NewRay(tuples.Point(0, 0, -5), tuples.Vector(0, 0, 1))
	// When
	xs := world.Intersect(*ray)
	// Then
	if len(*xs) != 4 {
		t.Errorf("Got %d intersections, expected %d", len(*xs), 4)
	}
	// And
	if (*xs)[0].Time != 4.0 {
		t.Errorf("xs[0].Time = %9.6f, expected %9.6f", (*xs)[0].Time, 4.0)
	}
	// And
	if (*xs)[1].Time != 4.5 {
		t.Errorf("xs[1].Time = %9.6f, expected %9.6f", (*xs)[1].Time, 4.5)
	}
	// And
	if (*xs)[2].Time != 5.5 {
		t.Errorf("xs[2].Time = %9.6f, expected %9.6f", (*xs)[2].Time, 5.5)
	}
	// And
	if (*xs)[3].Time != 6.0 {
		t.Errorf("xs[3].Time = %9.6f, expected %9.6f", (*xs)[3].Time, 6.0)
	}
}

// Scenario: Shading an intersection
// Given world ← default_world()
// And ray ← ray(point(0, 0, -5), vector(0, 0, 1))
// And shape ← the first object in world
// And hit ← intersection(4, shape)
// When prepare_hit(hit, ray)
// And c ← shade_hit(world, hit)
// Then c = color(0.38066, 0.47583, 0.2855)
func Test_Shading_an_Intersection(t *testing.T) {
	// Given
	world := DefaultWorld()
	// And
	ray := rays.NewRay(tuples.Point(0, 0, -5), tuples.Vector(0, 0, 1))
	// And
	shape := world.Objects[0]
	// And
	hit := rays.NewIntersection(4, &shape)
	// When
	hit.PrepareHit(*ray)
	// And
	c := world.ShadeHit(*hit)
	// Expected
	wantedC := colors.NewColor(0.38066, 0.47583, 0.2855)
	// Then
	if !wantedC.Equals(c) {
		t.Errorf("world.ShadeHit(%v) = %v, expected %v", hit, c, wantedC)
	}
}


// Scenario: Shading an intersection from the inside
// Given world ← default_world()
// And world.light ← point_light(point(0, 0.25, 0), color(1, 1, 1))
// And ray ← ray(point(0, 0, 0), vector(0, 0, 1))
// And shape ← the second object in world
// And hit ← intersection(0.5, shape)
// When prepare_hit(hit, ray)
// And c ← shade_hit(world, hit)
// Then c = color(0.90498, 0.90498, 0.90498)
func Test_Shading_an_Intersection_from_the_Inside(t *testing.T) {
	// Given
	world := DefaultWorld()
	// And
	light := lights.NewPointLight(tuples.Point(0, 0.25, 0), colors.NewColor(1, 1, 1))
	world.LightSource = &light
	// And
	ray := rays.NewRay(tuples.Point(0, 0, 0), tuples.Vector(0, 0, 1))
	// And
	shape := world.Objects[0]
	// And
	hit := rays.NewIntersection(4, &shape)
	// When
	hit.PrepareHit(*ray)
	// And
	c := world.ShadeHit(*hit)
	// Expected
	wantedC := colors.NewColor(0.90498, 0.90498, 0.90498)
	// Then
	if !wantedC.Equals(c) {
		t.Errorf("world.ShadeHit(%v) = %v, expected %v", hit, c, wantedC)
	}
}
