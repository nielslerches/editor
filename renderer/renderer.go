package renderer

import (
	"fmt"

	"github.com/nielslerches/editor/editor"
	"github.com/nielslerches/editor/line"
	"github.com/nielslerches/editor/screen"
	"github.com/nielslerches/editor/theme"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Renderer struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Theme    *theme.Theme
	Font     *ttf.Font
	Surface  *sdl.Surface
}

func NewRenderer(t *theme.Theme) (r *Renderer) {
	var err error
	var window *sdl.Window
	var renderer *sdl.Renderer
	var font *ttf.Font
	var surface *sdl.Surface
	var displayIndex int
	var ddpi, hdpi, vdpi float32

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	if window, err = sdl.CreateWindow(
		"Editor",
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		800, 600,
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE,
	); err != nil {
		panic(err)
	}

	if displayIndex, err = window.GetDisplayIndex(); err != nil {
		panic(err)
	}

	if ddpi, hdpi, vdpi, err = sdl.GetDisplayDPI(displayIndex); err != nil {
		panic(err)
	}

	fmt.Printf("ddpi, hdpi, vdpi: %v, %v, %v\n", ddpi, hdpi, vdpi)

	if surface, err = window.GetSurface(); err != nil {
		panic(err)
	}

	if renderer, err = sdl.CreateRenderer(
		window,
		-1,
		sdl.RENDERER_ACCELERATED,
	); err != nil {
		panic(err)
	}

	if err = ttf.Init(); err != nil {
		panic(err)
	}

	if font, err = ttf.OpenFont(
		t.Font.File,
		int(t.FontSize),
	); err != nil {
		panic(err)
	}

	font.SetKerning(true)

	fmt.Printf("Height: %v \n", font.Height())

	r = &Renderer{
		Window:   window,
		Renderer: renderer,
		Theme:    t,
		Font:     font,
		Surface:  surface,
	}

	return r
}

func (r *Renderer) Destroy() {
	var err error

	r.Font.Close()
	ttf.Quit()

	if err = r.Renderer.Destroy(); err != nil {
		panic(err)
	}

	if err = r.Window.Destroy(); err != nil {
		panic(err)
	}
}

func (r *Renderer) Resize(w, h int32) {
	fmt.Printf("Resize: %d, %d\n", w, h)
}

func (r *Renderer) Render(s *screen.Screen) {
	var err error

	if err = r.Renderer.SetDrawColor(r.Theme.Background.R, r.Theme.Background.G, r.Theme.Background.B, r.Theme.Background.A); err != nil {
		panic(err)
	}

	r.RenderScreen(s)
	r.Renderer.Present()

	fmt.Printf("Render\n")
}

func (r *Renderer) RenderScreen(s *screen.Screen) {
	r.RenderEditor(s.SelectedEditor)
}

func (r *Renderer) RenderEditor(e *editor.Editor) {
	if e == nil {
		return
	}

	for _, l := range e.Lines {
		index := e.GetLineIndex(l)

		if index < 0 {
			continue
		}

		number := uint(index)

		r.RenderLine(number, l)
	}
}

func (r *Renderer) RenderLine(number uint, l *line.Line) {
	if l == nil {
		return
	}

	var err error
	var surface *sdl.Surface
	var texture *sdl.Texture

	fmt.Printf("RenderLine: %d, %v", number, l)

	winWidth, _ := r.Window.GetSize()

	r.Renderer.SetDrawColor(r.Theme.Foreground.R, r.Theme.Foreground.G, r.Theme.Foreground.B, r.Theme.Foreground.A)

	rect := &sdl.Rect{
		X: 0,
		Y: int32(number * uint(r.Theme.LineHeight)),
		W: winWidth,
		H: int32(r.Theme.LineHeight),
	}

	if err = r.Renderer.FillRect(rect); err != nil {
		panic(err)
	}

	if len(l.Content) > 0 {
		var width, height int32
		color := sdl.Color{R: r.Theme.Text.R, G: r.Theme.Text.G, B: r.Theme.Text.B, A: r.Theme.Text.A}
		fg := sdl.Color{R: r.Theme.Foreground.R, G: r.Theme.Foreground.G, B: r.Theme.Foreground.B, A: r.Theme.Foreground.A}

		if surface, err = r.Font.RenderUTF8Shaded(
			l.Content, color, fg,
		); err != nil {
			panic(err)
		}

		defer surface.Free()

		if texture, err = r.Renderer.CreateTextureFromSurface(surface); err != nil {
			panic(err)
		}

		defer texture.Destroy()

		if _, _, width, height, err = texture.Query(); err != nil {
			panic(err)
		}

		fmt.Printf("Size: %v, %v\n", width, height)

		if err = r.Renderer.Copy(
			texture,
			nil,
			&sdl.Rect{
				X: rect.X,
				Y: rect.Y,
				W: width,
				H: height,
			},
		); err != nil {
			panic(err)
		}
	}
}
