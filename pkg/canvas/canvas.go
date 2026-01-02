package canvas

import "fmt"

type Canvas struct {
	width  int
	height int
	pixels [][]rune
}

func NewCanvas(width, height int) *Canvas {
	pixels := make([][]rune, height)

	for y := 0; y < height; y++ {
		pixels[y] = make([]rune, width)
		for x := 0; x < width; x++ {
			pixels[y][x] = '.'
		}
	}

	return &Canvas{
		width:  width,
		height: height,
		pixels: pixels,
	}
}

func (c *Canvas) DrawPixel(x, y int) {
	if x < 0 || x >= c.width || y < 0 || y >= c.height {
		return // fuera de pantalla
	}

	c.pixels[y][x] = '#'
}

func (c *Canvas) Clear() {
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			c.pixels[y][x] = '.'
		}
	}
}

func (c *Canvas) Present() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")

	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			fmt.Print(string(c.pixels[y][x]))
		}
		fmt.Println()
	}
}
