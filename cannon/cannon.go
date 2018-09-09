package main

import (
	"fmt"

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
	p := projectile{tuples.Point(0, 0, 10), tuples.Vector(1, 0, 0)}
	w := world{tuples.Vector(0, 0, -0.9), tuples.Vector(-0.5, 0, 0)}
	tick := 0

	for p.position.Z > 0 {
		tick = tick + 1
		p.position = p.position.Add(p.speed)
		p.speed = p.speed.Add(w.gravity).Add(w.wind)
		fmt.Printf("tick %d: position %v, speed %v\n", tick, p.position, p.speed)

	}
}
