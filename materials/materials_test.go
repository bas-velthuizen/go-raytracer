package materials

import (
	"github.com/bas-velthuizen/go-raytracer/colors"
	"github.com/bas-velthuizen/go-raytracer/tuples"
	"testing"
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
