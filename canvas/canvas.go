package canvas

import (
	"fmt"
	"github.com/bas-velthuizen/go-raytracer/colors"
	"log"
	"math"
	"os"
	"strconv"
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
	for y := 0; y < height; y++ {
		fmt.Printf("%d of %d\n", y, height)
		for x := 0; x < width; x++ {
			fmt.Print(".")
			c.grid[y*width+x] = colors.Color{Red: 0, Green: 0, Blue: 0}
		}
		fmt.Println()
	}
	fmt.Println()
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
	lines := []string{"P3", fmt.Sprintf("%d %d", c.Width, c.Height), "255"}
	for row := 0; row < c.Height; row++ {
		log.Printf("Row: %d of %d...", row+1, c.Height)
		stringList := make([]string, c.Width*c.Height*3)
		for col := 0; col < c.Width; col++ {
			color := c.Get(col, row)
			stringList = append(stringList, strconv.Itoa(toPPMColorComponent(color.Red)), strconv.Itoa(toPPMColorComponent(color.Green)), strconv.Itoa(toPPMColorComponent(color.Blue)))
		}
		line := ""
		for i:=0; i < len(stringList); i++ {
			if len(line + stringList[i]) >= 70 {
				lines = append(lines, line)
				line = ""
			}
			if len(line) > 0 {
				line += " "
			}
			line += stringList[i]
		}
		lines = append(lines, line)
	}
	log.Print("done")
	return PPM{lines}
}

// ToString writes the color pixmap to a string
func (p PPM) ToString() string {
	result := ""
	for i := 0; i < len(p.Lines); i++ {
		result += fmt.Sprintf("%s\n", p.Lines[i])
	}
	return result
}

// ToFile writes the color pixmap to a file
func (p PPM) ToFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 0; i < len(p.Lines); i++ {
		_, writeErr := file.WriteString( fmt.Sprintf("%s\n", p.Lines[i]) )
		if writeErr != nil {
			return err
		}
	}
	return nil
}

func toPPMColorComponent(component float64) int {
	value := int(math.Round(255 * component))
	if value < 0 {
		value = 0
	} else if value > 255 {
		value = 255
	}
	return value
}
