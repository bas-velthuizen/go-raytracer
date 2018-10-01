package main

import (
	"github.com/bas-velthuizen/go-raytracer/canvas"
	"github.com/bas-velthuizen/go-raytracer/colors"
	"github.com/bas-velthuizen/go-raytracer/rays"
	"github.com/bas-velthuizen/go-raytracer/transformations"
	"github.com/bas-velthuizen/go-raytracer/tuples"
)

func main() {
	red := colors.Color{Red: 1.0, Green: 0.0, Blue: 0.0}
	black := colors.Color{Red: 0.0, Green: 0.0, Blue: 0.0}

	c := canvas.NewCanvas(250, 250)

	scale := transformations.Scaling(0.5, 1, 1)
	shear := transformations.Shearing(1, 0, 0, 0, 0, 0)
	// rot := transformations.RotationZ(-math.Pi / 3.0)
	trans := transformations.Translation(0, 0, 2)

	origin := tuples.Point(0, 0, -5)
	s := rays.NewUnitSphere()
	s.SetTransform(trans.Multiply(*scale).Multiply(*shear))
	depth := 10.0
	width := 7.0
	step := width / float64(c.Width)

	for row := 0; row < c.Width; row++ {
		y := width/2.0 - float64(row)*step
		for col := 0; col < c.Width; col++ {
			x := float64(col)*step - width/2.0
			target := tuples.Point(x, y, depth)
			direction := target.Subtract(origin)
			ray := rays.NewRay(origin, direction)
			xs := ray.Intersect(s)
			if xs.Hit() != nil {
				c.Set(col, row, red)
			} else {
				c.Set(col, row, black)
			}
		}
	}

	c.ToPPM().ToFile("picture.ppm")
}
