package canvas

import (
	"fmt"
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
func Test_Write_to_Canvas(t *testing.T) {
	// Given
	c := NewCanvas(10, 20)
	red := colors.Color{Red: 1, Green: 0, Blue: 0}
	// When
	c.Set(2, 3, red)
	// then
	pixel := c.Get(2,3)
	if !red.Equals(pixel) {
		t.Errorf("c.Pixel(%d,%d) == %v, want %v", 2, 3, pixel, red)
	}
}

// Scenario: Constructing the PPM header
// Given c ← canvas(5, 3)
// When ppm ← canvas_to_ppm(c)
// Then lines 1-3 of ppm are
// """
// P3
// 5 3
// 255
// """
func Test_Construct_PPM_Header(t *testing.T) {
	// Given
	c:= NewCanvas(5, 3)
	// When
	ppm := c.ToPPM()
	// Expected
	lines :=   []string{"P3","5 3","255"}
	// Then
	for i := 0; i < 3; i++ {
		if ppm.Lines[i] != lines[i] {
			t.Errorf("line[%d] == %s, want %s", i, ppm.Lines[i], lines[i])
		}
	}
}

// Scenario: Constructing the PPM pixel data
// Given c ← canvas(5, 3)
// And c1 ← color(1.5, 0, 0)
// And c2 ← color(0, 0.5, 0)
// And c3 ← color(-0.5, 0, 1)
// When write_pixel(c, 0, 0, c1)
// And write_pixel(c, 2, 1, c2)
// And write_pixel(c, 4, 2, c3)
// And ppm ← canvas_to_ppm(c)
// Then lines 4-6 of ppm are
// """
// 255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
// 0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
// 0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
// """
func Test_Construct_PPM_Pixel_Data(t *testing.T) {
	// Given
	c:= NewCanvas(5, 3)
	c1 := colors.Color{ Red: 1.5,  Green: 0,   Blue: 0}
	c2 := colors.Color{ Red: 0,    Green: 0.5, Blue: 0}
	c3 := colors.Color{ Red: -0.5, Green: 0,   Blue: 1}
	// When
	c.Set(0, 0, c1)
	c.Set(2, 1, c2)
	c.Set(4, 2, c3)
	ppm := c.ToPPM()
	// Expected
	lines := []string{
		"255 0 0 0 0 0 0 0 0 0 0 0 0 0 0",
		"0 0 0 0 0 0 0 128 0 0 0 0 0 0 0",
		"0 0 0 0 0 0 0 0 0 0 0 0 0 0 255",
	}
	// Then
	for i := 3; i < 6; i++ {
		if ppm.Lines[i] != lines[i - 3] {
			t.Errorf("line[%d] == %s, want %s", i, ppm.Lines[i], lines[i-3])
		}
	}
}

// Scenario: Splitting long lines in PPM files
// Given c ← canvas(10, 2)
// When every pixel of c is set to color(1, 0.8, 0.6)
// And ppm ← canvas_to_ppm(c)
// Then lines 4-7 of ppm are
// """
// 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
// 153 255 204 153 255 204 153 255 204 153 255 204 153
// 255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
// 153 255 204 153 255 204 153 255 204 153 255 204 153
// """
func Test_Construct_PPM_Pixel_Data_with_Long_Lines(t *testing.T) {
	// Given
	c := NewCanvas(10, 2)
	//When
	for y:=0; y < 2; y++ {
		for x:=0; x < 10; x++ {
			c.Set(x, y, colors.Color{Red:1, Green: 0.8, Blue: 0.6})
		}
	}
	ppm := c.ToPPM()
	// Expected
	lines := []string{
		"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204",
		"153 255 204 153 255 204 153 255 204 153 255 204 153",
		"255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204",
		"153 255 204 153 255 204 153 255 204 153 255 204 153",
	}
	// Then
	for i := 3; i < 7; i++ {
		if ppm.Lines[i] != lines[i - 3] {
			t.Errorf("line[%d] == %s (%d), want %s (%d)", i, ppm.Lines[i], len(ppm.Lines[i]), lines[i-3], len(lines[i-3]))
		}
	}
}

// Scenario: PPM files are terminated by a newline
// Given c ← canvas(5, 3)
// When ppm ← canvas_to_ppm(c)
// Then the last character of ppm is a newline
func Test_PPM_Terminated_by_Newline(t *testing.T) {
	// Given
	c := NewCanvas(5, 3)
	// When
	ppm := c.ToPPM().ToString()
	// Then
	if ppm[len(ppm)-1] != '\n' {
		t.Errorf("Expected %s as last character, was %s", fmt.Sprintf("% x", '\n'), fmt.Sprintf("% x", ppm[len(ppm)-1]))
	}

}