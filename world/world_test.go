package world

import (
	"testing"

	"github.com/bas-velthuizen/go-raytracer/colors"
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
	w := NewWorld([]rays.Sphere{}, nil)
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
	s1 := rays.NewUnitSphere()
	s1.Material.Color = colors.NewColor(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2
	// And
	s2 := rays.NewUnitSphere()
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
