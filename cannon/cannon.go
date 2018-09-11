package main

import (
	"fmt"
	"github.com/bas-velthuizen/go-raytracer/canvas"
	"github.com/bas-velthuizen/go-raytracer/colors"
	"math"

	"github.com/bas-velthuizen/go-raytracer/tuples"
)

type projectile struct {
	position tuples.Tuple
	speed    tuples.Tuple
}

type world struct {
	wind    tuples.Tuple
	gravity tuples.Tuple
}

func main() {
	start := tuples.Point(0,1,0)
	velocity := tuples.Vector(1,1.8,0).Normalize().Multiply(11.25 )
	p := projectile{start, velocity}

	wind := tuples.Vector(-0.01, 0, 0)
	gravity := tuples.Vector(0, -0.1, 0)
	w := world{wind, gravity}

	tick := 0

	c := canvas.NewCanvas(900, 550)
	red := colors.Color{Red: 1, Green: 0, Blue: 0}

	fmt.Println("Canvas created, now shooting")
	for p.position.Y > 0 {
		x := int(math.Round(p.position.X))
		y := 550 - int(math.Round(p.position.Y))
		c.Set(x, y, red)

		tick = tick + 1
		p.position = p.position.Add(p.speed)
		p.speed = p.speed.Add(w.gravity).Add(w.wind)
		fmt.Printf("tick %d: position %v, speed %v\n", tick, p.position, p.speed)
	}

	fmt.Println("Curve finished, now saving")
	err := c.ToPPM().ToFile("./curve.ppm")
	if err != nil {
		fmt.Println(err)
	}
}
