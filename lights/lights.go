package lights

import (
	"fmt"
	"github.com/bas-velthuizen/go-raytracer/colors"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

type PointLight struct {
	Position tuples.Tuple
	Intensity colors.Color
}

// NewPointLight constructs a new Point Light from a Point and a Color
func NewPointLight(position tuples.Tuple, intensity colors.Color) PointLight {
	return PointLight{position, intensity}
}

func (p PointLight) String() string {
	return fmt.Sprintf("PointLight( %v, %v )", p.Position, p.Intensity)
}