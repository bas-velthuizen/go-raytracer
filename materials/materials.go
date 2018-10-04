package materials

import (
	"fmt"
	"github.com/bas-velthuizen/go-raytracer/colors"
	"github.com/bas-velthuizen/go-raytracer/tuples"
	"math"
)

// Materual defines the properties of a material
type Material struct {
	Color     colors.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

// DefaultMaterial constructs the default material
func DefaultMaterial() Material {
	return Material{
		colors.NewColor(1, 1, 1),
		0.1,
		0.9,
		0.9,
		200.0,
	}
}

func (m Material) Equals(other Material) bool {
	return m.Color.Equals(other.Color) &&
		math.Abs(m.Ambient - other.Ambient) <= tuples.Epsilon &&
		math.Abs(m.Diffuse - other.Diffuse) <= tuples.Epsilon &&
		math.Abs(m.Specular - other.Specular) <= tuples.Epsilon &&
		math.Abs(m.Shininess - other.Shininess) <= tuples.Epsilon
}

func (m Material) String() string {
	return fmt.Sprintf("Material( %v, %9.6f, %9.6f, %9.6f, %9.6f )", m.Color, m.Ambient, m.Diffuse, m.Specular, m.Shininess)
}
