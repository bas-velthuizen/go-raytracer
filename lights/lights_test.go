package lights

import (
	"testing"

	"github.com/bas-velthuizen/go-raytracer/colors"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// Scenario: A point light has a position and intensity
// Given intensity ← color(1, 1, 1)
// And position ← point(0, 0, 0)
// When light ← point_light(position, intensity)
// Then light.position = position
// And light.intensity = intensity
func Test_a_Point_Light_has_a_Position_and_Intensity(t *testing.T) {
	// Given
	intensity := colors.White()
	// And
	position := tuples.Point(0, 0, 0)
	// When
	light := NewPointLight(position, intensity)
	// Then
	if !position.Equals(light.Position) {
		t.Errorf("%v has position %v, expected %v", light, light.Position, position)
	}
	// And
	if !intensity.Equals(light.Intensity) {
		t.Errorf("%v has intensity %v, expected %v", light, light.Intensity, intensity)
	}
}
