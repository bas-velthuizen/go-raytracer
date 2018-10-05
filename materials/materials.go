package materials

import (
	"fmt"
	"math"

	"github.com/bas-velthuizen/go-raytracer/lights"

	"github.com/bas-velthuizen/go-raytracer/colors"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// Material defines the properties of a material
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
		colors.White(),
		0.1,
		0.9,
		0.9,
		200.0,
	}
}

// Equals checks if this material is the same as another
func (m Material) Equals(other Material) bool {
	return m.Color.Equals(other.Color) &&
		math.Abs(m.Ambient-other.Ambient) <= tuples.Epsilon &&
		math.Abs(m.Diffuse-other.Diffuse) <= tuples.Epsilon &&
		math.Abs(m.Specular-other.Specular) <= tuples.Epsilon &&
		math.Abs(m.Shininess-other.Shininess) <= tuples.Epsilon
}

func (m Material) String() string {
	return fmt.Sprintf("Material( %v, %9.6f, %9.6f, %9.6f, %9.6f )", m.Color, m.Ambient, m.Diffuse, m.Specular, m.Shininess)
}

// Lighting calculates the effective color of a pixel with reflections of light
func (m Material) Lighting(
	light lights.PointLight,
	position tuples.Tuple,
	eyeV tuples.Tuple,
	normalV tuples.Tuple,
) colors.Color {
	diff := colors.Black()
	spec := colors.Black()

	effectiveColor := m.Color.Blend(light.Intensity)
	lightV := light.Position.Subtract(position).Normalize()

	ambient := effectiveColor.Multiply(m.Ambient)

	lightDotNormal := lightV.Dot(normalV)

	if lightDotNormal < 0 {
		diff = colors.Black()
		spec = colors.Black()
	} else {
		diff = effectiveColor.Multiply(m.Diffuse * lightDotNormal)
		reflectV := normalV.Reflect(lightV.Negate())
		reflectDotEye := math.Pow(reflectV.Dot(eyeV), m.Shininess)
		if reflectDotEye <= 0 {
			spec = colors.Black()
		} else {
			spec = light.Intensity.Multiply(m.Specular * reflectDotEye)
		}
	}
	return ambient.Add(diff).Add(spec)

}
