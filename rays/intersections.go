package rays

import "fmt"

// Intersection aggregates a time value and a Sphere
type Intersection struct {
	Time   float64
	Object *Sphere
}

// Intersections groups multiple Intersection instances
type Intersections []*Intersection

// NewIntersection creates a new intersection and passes a reference to it
func NewIntersection(time float64, object *Sphere) *Intersection {
	return &Intersection{Time: time, Object: object}
}

// String formats Intersection to readable string
func (i Intersection) String() string {
	return fmt.Sprintf("Intersection( %9.6f, %v )", i.Time, i.Object)
}

// NewIntersections creates a new collection of Intersection references
func NewIntersections(is []*Intersection) *Intersections {
	result := make(Intersections, len(is))
	copy(result, is)
	return &result
}

// Hit calculates the hit closest to the view point
func (xs Intersections) Hit() *Intersection {
	if len(xs) == 0 {
		return nil
	}
	minIndex := -1
	for i := 0; i < len(xs); i++ {
		x := xs[i]
		if x.Time > 0 {
			if minIndex < 0 {
				minIndex = i
			} else {
				if xs[minIndex].Time > xs[i].Time {
					minIndex = i
				}
			}
		}
	}
	if minIndex < 0 {
		return nil
	}
	return xs[minIndex]
}
