package canvas

import (
	"fmt"
	"github.com/bas-velthuizen/go-raytracer/colors"
	"math"
)

// Canvas implements a canvas on which bitmap images can be projected/drawn
type Canvas struct {
	Width  int
	Height int
	grid   []colors.Color
}

type PPM struct {
	Lines []string
}

// NewCanvas creates a new Canvas and returns a pointer to it
func NewCanvas(width int, height int) *Canvas {
	c := Canvas{width, height, make([]colors.Color, width*height)}
	for i := 0; i < cap(c.grid); i++ {
		c.grid[i] = colors.Color{Red: 0, Green: 0, Blue: 0}
	}
	return &c
}

// Get returns the value of the pixel at the specified position
func (c Canvas) Get(x int, y int) colors.Color {
	return c.grid[y*c.Width+x]
}

// Set sets the value of the pixel at the specified position
func (c Canvas) Set(x int, y int, color colors.Color) {
	c.grid[y*c.Width+x] = color
}

// ToPPM creates a PPM data structure of the Canvas
func (c Canvas) ToPPM() PPM {
	lines := []string{"P3", "5 3", "255"}
	for row := 0; row < c.Height; row++ {
		line := ""
		for col := 0; col < c.Width; col++ {
			color := c.Get(col, row)
			toPPMColorComponent(color.Red)

			line += fmt.Sprintf("%d %d %d", toPPMColorComponent(color.Red), toPPMColorComponent(color.Green), toPPMColorComponent(color.Blue))
			if col != c.Width-1 {
				line += " "
			}
		}
		lines = append(lines, line)
	}
	return PPM{lines}
}

func toPPMColorComponent(component float64) uint {
	value := int(math.Round(255 * component))
	if value < 0 {
		value = 0
	} else if value > 255 {
		value = 255
	}
	return uint(value)
}
