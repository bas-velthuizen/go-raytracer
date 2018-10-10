package world

import (
	"sort"

	"github.com/bas-velthuizen/go-raytracer/colors"
	"github.com/bas-velthuizen/go-raytracer/lights"
	"github.com/bas-velthuizen/go-raytracer/rays"
	"github.com/bas-velthuizen/go-raytracer/spheres"
	"github.com/bas-velthuizen/go-raytracer/transformations"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

// World defines the light sources and objects in a world
type World struct {
	Objects      []spheres.Sphere
	LightSources []lights.PointLight
}

// NewWorld returns a new World object with the provides Objects and Light Source
func NewWorld(Objects []spheres.Sphere, LightSources []lights.PointLight) World {
	return World{Objects, LightSources}
}

// DefaultWorld returns a new Default World object
func DefaultWorld() World {
	light := lights.NewPointLight(tuples.Point(-10, 10, -10), colors.White())

	s1 := spheres.NewUnitSphere()
	s1.Material.Color = colors.NewColor(0.8, 1.0, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2

	s2 := spheres.NewUnitSphere()
	s2.SetTransform(transformations.Scaling(0.5, 0.5, 0.5))

	return World{Objects: []spheres.Sphere{*s1, *s2}, LightSources: []lights.PointLight{light}}
}

// Contains checks whether the world contains this object
func (w World) Contains(object spheres.Sphere) bool {
	for i := 0; i < len(w.Objects); i++ {
		if object.Equals(w.Objects[i]) {
			return true
		}
	}
	return false
}

// Intersect calculates the intersections with all objects in the world
func (w World) Intersect(ray rays.Ray) *rays.Intersections {
	xsArray := make([]*rays.Intersection, 0, 0)
	for i := 0; i < len(w.Objects); i++ {
		partXsArray := ray.Intersect(&w.Objects[i])
		xsArray = append(xsArray, partXsArray...)
	}
	sort.Sort(rays.ByTime(xsArray))
	return rays.NewIntersections(xsArray)
}

// ShadeHit calculates the color of a hit in the world
func (w World) ShadeHit(hit rays.Intersection) colors.Color {
	result := colors.Black()
	for i := 0; i < len(w.LightSources); i++ {
		c := hit.Object.Material.Lighting(
			w.LightSources[i],
			hit.Point,
			hit.EyeV,
			hit.NormalV)
		result = result.Add(c)
	}
	return result
}

// ColorAt calculates the color caused by a ray
func (w World) ColorAt(ray rays.Ray) colors.Color {
	// 1. Call intersect_world to find the intersections of the given ray with the given
	// world.
	xs := w.Intersect(ray)
	// 2. Find the hit from the resulting intersections.
	hit := xs.Hit()
	// 3. Return the color black if there is no hit.
	if hit == nil {
		return colors.Black()
	}
	// 4. Otherwise, prepare the hit with prepare_hit.
	hit.PrepareHit(ray)
	// 5. Finally, call shade_hit to find the color at the hit intersection.
	result := w.ShadeHit(*hit)
	return result
}
