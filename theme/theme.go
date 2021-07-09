package theme

import "github.com/nielslerches/editor/fonts"

type Theme struct {
	Background *Color
	Foreground *Color
	LineHeight uint8
	FontSize   int
	Font       *fonts.Font
	Text       *Color
}

func NewTheme(background, foreground *Color, lineHeight uint8, fontSize int, font *fonts.Font, text *Color) (t *Theme) {
	t = &Theme{
		Background: background,
		Foreground: foreground,
		LineHeight: lineHeight,
		Font:       font,
		Text:       text,
	}

	return t
}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func NewColor(r, g, b, a uint8) (c *Color) {
	c = &Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}

	return c
}

func (c *Color) RGBA() (uint8, uint8, uint8, uint8) {
	return c.R, c.G, c.B, c.A
}
