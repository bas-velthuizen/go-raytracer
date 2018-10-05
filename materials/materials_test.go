package materials

import (
	"math"
	"testing"

	"github.com/bas-velthuizen/go-raytracer/colors"
	"github.com/bas-velthuizen/go-raytracer/lights"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// setup returns the default values for material and position
func setup() (Material, tuples.Tuple) {
	return DefaultMaterial(), tuples.Point(0, 0, 0)
}

// Scenario: The default material
// Given m ← material()
// Then m.color = color(1, 1, 1)
// And m.ambient = 0.1
// And m.diffuse = 0.9
// And m.specular = 0.9
// And m.shininess = 200
func Test_the_Default_Material(t *testing.T) {
	// Given
	m := DefaultMaterial()
	// Then

	wantedColor := colors.NewColor(1, 1, 1)
	if !wantedColor.Equals(m.Color) {
		t.Errorf("%v has color %v, expected %v", m, m.Color, wantedColor)
	}
	// And
	wantedAmbient := 0.1
	if wantedAmbient != m.Ambient {
		t.Errorf("%v has ambient %9.6f, expected %9.6f", m, m.Ambient, wantedAmbient)
	}
	// And
	wantedDiffuse := 0.9
	if wantedDiffuse != m.Diffuse {
		t.Errorf("%v has ambient %9.6f, expected %9.6f", m, m.Diffuse, wantedDiffuse)
	}
	// And
	wantedSpecular := 0.9
	if wantedSpecular != m.Specular {
		t.Errorf("%v has ambient %9.6f, expected %9.6f", m, m.Specular, wantedSpecular)
	}
	// And
	wantedShininess := 200.0
	if wantedShininess != m.Shininess {
		t.Errorf("%v has ambient %9.6f, expected %9.6f", m, m.Shininess, wantedShininess)
	}
}

// Scenario: Lighting with the eye between the light and the surface
// Given eyev ← vector(0, 0, -1)
// And normalv ← vector(0, 0, -1)
// And light ← point_light(point(0, 0, -10), color(1, 1, 1))
// When result ← lighting(m, light, position, eyev, normalv)
// Then result = color(1.9, 1.9, 1.9)
func Test_Lighting_With_the_Eye_Between_the_Light_and_the_Surface(t *testing.T) {
	// Setup
	m, position := setup()
	// Given
	eyev := tuples.Vector(0, 0, -1)
	// And
	normalv := tuples.Vector(0, 0, -1)
	// And
	light := lights.NewPointLight(tuples.Point(0, 0, -10), colors.NewColor(1, 1, 1))
	// When
	result := m.Lighting(light, position, eyev, normalv)
	// Expected
	wanted := colors.NewColor(1.9, 1.9, 1.9)
	// Then
	if !wanted.Equals(result) {
		t.Errorf("Lighting( %v, %v, %v, %v ) = %v, Expected %v", m, light, eyev, normalv, result, wanted)
	}
}

// Scenario: Lighting with the eye between light and surface, eye offset 45°
// Given eyev ← vector(0, √2/2, -√2/2)
// And normalv ← vector(0, 0, -1)
// And light ← point_light(point(0, 0, -10), color(1, 1, 1))
// When result ← lighting(m, light, position, eyev, normalv)
// Then result = color(1.0, 1.0, 1.0)
func Test_Lighting_With_the_Eye_Between_Light_and_Surface__Eye_Offset_45(t *testing.T) {
	// Setup
	m, position := setup()
	// Given
	eyev := tuples.Vector(0, math.Sqrt2/2.0, -math.Sqrt2/2)
	// And
	normalv := tuples.Vector(0, 0, -1)
	// And
	light := lights.NewPointLight(tuples.Point(0, 0, -10), colors.NewColor(1, 1, 1))
	// When
	result := m.Lighting(light, position, eyev, normalv)
	// Expected
	wanted := colors.NewColor(1.0, 1.0, 1.0)
	// Then
	if !wanted.Equals(result) {
		t.Errorf("Lighting( %v, %v, %v, %v ) = %v, Expected %v", m, light, eyev, normalv, result, wanted)
	}
}

// Scenario: Lighting with eye opposite surface, light offset 45°
// Given eyev ← vector(0, 0, -1)
// And normalv ← vector(0, 0, -1)
// And light ← point_light(point(0, 10, -10), color(1, 1, 1))
// When result ← lighting(m, light, position, eyev, normalv)
// Then result = color(0.7364, 0.7364, 0.7364)
func Test_Lighting_With_Eye_Opposite_Surface__Light_Offset_45(t *testing.T) {
	// Setup
	m, position := setup()
	// Given
	eyev := tuples.Vector(0, 0, -1)
	// And
	normalv := tuples.Vector(0, 0, -1)
	// And
	light := lights.NewPointLight(tuples.Point(0, 10, -10), colors.NewColor(1, 1, 1))
	// When
	result := m.Lighting(light, position, eyev, normalv)
	// Expected
	wanted := colors.NewColor(0.7364, 0.7364, 0.7364)
	// Then
	if !wanted.Equals(result) {
		t.Errorf("Lighting( %v, %v, %v, %v ) = %v, Expected %v", m, light, eyev, normalv, result, wanted)
	}
}

// Scenario: Lighting with eye in the path of the reflection vector
// Given eyev ← vector(0, -√2/2, -√2/2)
// And normalv ← vector(0, 0, -1)
// And light ← point_light(point(0, 10, -10), color(1, 1, 1))
// When result ← lighting(m, light, position, eyev, normalv)
// Then result = color(1.6364, 1.6364, 1.6364)
func Test_Lighting_With_Eye_in_the_Path_of_the_Reflection_Vector(t *testing.T) {
	// Setup
	m, position := setup()
	// Given
	eyev := tuples.Vector(0, -math.Sqrt2/2.0, -math.Sqrt2/2.0)
	// And
	normalv := tuples.Vector(0, 0, -1)
	// And
	light := lights.NewPointLight(tuples.Point(0, 10, -10), colors.NewColor(1, 1, 1))
	// When
	result := m.Lighting(light, position, eyev, normalv)
	// Expected
	wanted := colors.NewColor(1.6364, 1.6364, 1.6364)
	// Then
	if !wanted.Equals(result) {
		t.Errorf("Lighting( %v, %v, %v, %v ) = %v, Expected %v", m, light, eyev, normalv, result, wanted)
	}
}

// Scenario: Lighting with the light behind the surface
// Given eyev ← vector(0, 0, -1)
// And normalv ← vector(0, 0, -1)
// And light ← point_light(point(0, 0, 10), color(1, 1, 1))
// When result ← lighting(m, light, position, eyev, normalv)
// Then result = color(0.1, 0.1, 0.1)
func Test_Lighting_with_the_Light_Behind_the_Surface(t *testing.T) {
	// Setup
	m, position := setup()
	// Given
	eyev := tuples.Vector(0, 0, -1)
	// And
	normalv := tuples.Vector(0, 0, -1)
	// And
	light := lights.NewPointLight(tuples.Point(0, 0, 10), colors.NewColor(1, 1, 1))
	// When
	result := m.Lighting(light, position, eyev, normalv)
	// Expected
	wanted := colors.NewColor(0.1, 0.1, 0.1)
	// Then
	if !wanted.Equals(result) {
		t.Errorf("Lighting( %v, %v, %v, %v ) = %v, Expected %v", m, light, eyev, normalv, result, wanted)
	}
}
