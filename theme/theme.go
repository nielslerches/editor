package theme

import "image/color"

type Theme struct {
	Background color.RGBA
	Foreground color.RGBA
	LineHeight uint8
}

var (
	DEFAULT = &Theme{
		Background: color.RGBA{R: 127, G: 127, B: 127, A: 255},
		Foreground: color.RGBA{R: 63, G: 63, B: 63, A: 255},
		LineHeight: 32,
	}
)
