package canvas

import (
	"github.com/bas-velthuizen/go-raytracer/colors"
)

// Canvas implements a canvas on which bitmap images can be projected/drawn
type Canvas struct {
	Width  int
	Height int
	grid   []colors.Color
}

// NewCanvas creates a new Canvas and returns a pointer to it
func NewCanvas(width int, height int) *Canvas {
	c := Canvas{width, height, make([]colors.Color, width*height)}
	for i := 0; i < cap(c.grid); i++ {
		c.grid[i] = colors.Color{Red: 0, Green: 0, Blue: 0}
	}
	return &c
}

// Pixel returns the value of the pixel at the specified position
func (c Canvas) Get(x int, y int) colors.Color {
	return c.grid[y*c.Width+x]
}
