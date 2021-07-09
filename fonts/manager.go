package fonts

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/flopp/go-findfont"
	"golang.org/x/image/font/sfnt"
)

type FontManager struct {
	Fonts  []*Font
	Loaded bool
}

func NewFontManager() (m *FontManager) {
	m = &FontManager{}
	m.LoadFonts()

	return m
}

func (m *FontManager) GetFonts() (fs *FontSet) {
	if !m.Loaded {
		m.LoadFonts()
	}

	fs = &FontSet{
		Result: m.Fonts,
	}

	return fs
}

func (m *FontManager) LoadFonts() {
	files := findfont.List()
	fonts := make([]*Font, 0)

	b := &sfnt.Buffer{}

	for _, file := range files {
		var err error
		var data []byte
		var ttf *sfnt.Font
		var family string
		var style string

		if data, err = ioutil.ReadFile(file); err != nil {
			continue
		}

		if ttf, err = sfnt.Parse(data); err != nil {
			continue
		}

		if family, err = ttf.Name(b, sfnt.NameIDFamily); err != nil {
			panic(err)
		}

		if strings.Contains(strings.ToLower(family), "mono") {
			style = "monospace"
		} else if strings.Contains(strings.ToLower(family), "sans") {
			style = "sans-serif"
		} else if strings.Contains(strings.ToLower(family), "serif") {
			style = "serif"
		}

		units := ttf.UnitsPerEm()

		fmt.Printf("UnitsPerEm: %v\n", units)

		fonts = append(fonts, &Font{
			File:   file,
			Family: family,
			Style:  style,
			Units:  int32(units),
		})
	}

	m.Fonts = fonts
}
