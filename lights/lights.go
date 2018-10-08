package lights

import (
	"fmt"

	"github.com/bas-velthuizen/go-raytracer/colors"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// PointLight defines a light source from a single point with a certain intensity and color
type PointLight struct {
	Position  tuples.Tuple
	Intensity colors.Color
}

// NewPointLight constructs a new Point Light from a Point and a Color
func NewPointLight(position tuples.Tuple, intensity colors.Color) PointLight {
	return PointLight{position, intensity}
}

func (p PointLight) String() string {
	return fmt.Sprintf("PointLight( %v, %v )", p.Position, p.Intensity)
}

// Equals checks if another PointLight is the same as the current one
func (p PointLight) Equals(other PointLight) bool {
	return p.Position.Equals(other.Position) && p.Intensity.Equals(other.Intensity)
}
