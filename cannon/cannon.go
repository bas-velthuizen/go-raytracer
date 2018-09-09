package main

import (
	"fmt"
)

type projectile struct {
	position Tuple
	speed    Tuple
}

type world struct {
	wind    Tuple
	gravity Tuple
}

func main() {
	p := projectile{Point(0, 0, 10), Vector(1, 0, 0)}
	w := world{Vector(0, 0, -0.1), Vector(0, 0, 0)}
	tick := 0

	for p.Z > 0 {
		tick = tick + 1
		p.position = p.position.Add(p.speed)
		p.speed = p.speed.Add(w.gravity).Add(w.wind)
		fmt.Printf("tick %d: position %v, speed %v\n", tick, p.position, p.speed)

	}
}
