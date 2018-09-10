package canvas

import (
	"testing"

	"github.com/bas-velthuizen/go-raytracer/colors"
)

// Scenario: Creating a canvas
// Given c ← canvas(10, 20)
// Then c.width = 10
// And c.height = 20
// And every pixel of c is color(0, 0, 0)
func Test_Create_Canvas(t *testing.T) {
	// Given
	c := NewCanvas(10, 20)
	// Then
	if c.Width != 10 {
		t.Errorf("c.Width == %d want %d", c.Width, 10)
	}
	if c.Height != 20 {
		t.Errorf("c.Height == %d, want %d", c.Height, 10)
	}
	blank := colors.Color{Red: 0, Green: 0, Blue: 0}
	for x := 0; x < c.Width; x++ {
		for y := 0; y < c.Height; y++ {
			pixel := c.Get(x, y)
			if !blank.Equals(pixel) {
				t.Errorf("c.Pixel(%d,%d) == %v, want %v", x, y, pixel, blank)
			}
		}
	}
}

// Scenario: Writing pixels to a canvas
// Given c ← canvas(10, 20)
// And red ← color(1, 0, 0)
// When write_pixel(c, 2, 3, red)
// Then pixel_at(c, 2, 3) = red
