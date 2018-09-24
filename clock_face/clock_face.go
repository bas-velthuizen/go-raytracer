package main

import (
	"fmt"
	"github.com/bas-velthuizen/go-raytracer/canvas"
	"github.com/bas-velthuizen/go-raytracer/colors"
	"github.com/bas-velthuizen/go-raytracer/transformations"
	"github.com/bas-velthuizen/go-raytracer/tuples"
	"math"
)

func main() {
	c := canvas.NewCanvas(100, 100)
	center := transformations.Translation(50, 50, 0)
	scale := transformations.Scaling(40, 40, 0)
	white := colors.Color{Red: 1.0, Green: 1.0, Blue: 1.0}
	p := tuples.Point(0, 1, 0 )
	for i := 0; i < 12; i++ {
		angle := math.Pi * float64(i) / 6
		rotate := transformations.RotationZ(angle)
		rp := center.Multiply(*scale).Multiply(*rotate).MultiplyTuple(p)
		x := int(math.Round(rp.X))
		y := int(math.Round(rp.Y))
		fmt.Printf("center()*scale(40)*rotate( %9.6f )* %v => %v (%d, %d)\n", angle, p, rp, x, y)
		c.Set(x, y, white)
	}

	c.ToPPM().ToFile("clockface.ppm")
}
