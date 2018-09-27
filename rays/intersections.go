package rays

import "fmt"

// Intersection aggregates a time value and a Sphere
type Intersection struct {
	time float64
	object *Sphere
}

// Intersections groups multiple Intersection instances
type Intersections []*Intersection

// NewIntersection creates a new intersection and passes a reference to it
func NewIntersection(time float64, object *Sphere) *Intersection {
	return &Intersection{time: time, object: object}
}

// String formats Intersection to readable string
func (i Intersection) String() string {
	return fmt.Sprintf("Intersection( %9.6f, %v )", i.time, i.object)
}

// NewIntersections creates a new collection of Intersection references
func NewIntersections(is []*Intersection) *Intersections {
	result := make(Intersections, len(is))
	copy( result, is)
	return &result
}
