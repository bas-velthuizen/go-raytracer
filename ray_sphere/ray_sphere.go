package main

import (
	"github.com/bas-velthuizen/go-raytracer/canvas"
	"github.com/bas-velthuizen/go-raytracer/colors"
	"github.com/bas-velthuizen/go-raytracer/lights"
	"github.com/bas-velthuizen/go-raytracer/rays"
	"github.com/bas-velthuizen/go-raytracer/transformations"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

func main() {
	black := colors.Black()

	c := canvas.NewCanvas(250, 250)

	scale := transformations.Scaling(1, 1, 1)
	// shear := transformations.Shearing(1, 0, 0, 0, 0, 0)
	// rot := transformations.RotationY(-math.Pi / 4.0)
	trans := transformations.Translation(0, 0, 2)

	origin := tuples.Point(0, 0, -5)
	s := rays.NewUnitSphere()
	s.Material.Color = colors.NewColor(0.2, 1, 1)
	s.SetTransform(trans.Multiply(*scale) /*.Multiply(*rot).Multiply(*shear)*/)
	depth := 10.0
	width := 7.0
	step := width / float64(c.Width)

	// light source
	lightPosition := tuples.Point(-10, 10, -10)
	lightColor := colors.White()
	light := lights.NewPointLight(lightPosition, lightColor)

	for row := 0; row < c.Width; row++ {
		y := width/2.0 - float64(row)*step
		for col := 0; col < c.Width; col++ {
			x := float64(col)*step - width/2.0
			target := tuples.Point(x, y, depth)
			direction := target.Subtract(origin).Normalize()
			ray := rays.NewRay(origin, direction)
			xs := ray.Intersect(s)
			if xs.Hit() != nil {
				point := ray.Position(xs.Hit().Time)
				normal := xs.Hit().Object.NormalAt(*point)
				eye := ray.Direction.Negate()
				color := xs.Hit().Object.Material.Lighting(light, *point, eye, *normal)
				c.Set(col, row, color)
			} else {
				c.Set(col, row, black)
			}
		}
	}

	c.ToPPM().ToFile("picture.ppm")
}
