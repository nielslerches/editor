package renderer

import (
	"fmt"

	"github.com/nielslerches/editor/editor"
	"github.com/nielslerches/editor/line"
	"github.com/nielslerches/editor/screen"
	"github.com/nielslerches/editor/theme"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type Renderer struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Theme    *theme.Theme
}

func NewRenderer(t *theme.Theme) (r *Renderer) {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error

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

	if renderer, err = sdl.CreateRenderer(
		window,
		-1,
		sdl.RENDERER_ACCELERATED,
	); err != nil {
		panic(err)
	}

	r = &Renderer{
		Window:   window,
		Renderer: renderer,
		Theme:    t,
	}

	return r
}

func (r *Renderer) Destroy() {
	var err error

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
	r.Renderer.SetDrawColor(r.Theme.Background.R, r.Theme.Background.G, r.Theme.Background.B, r.Theme.Background.A)
	if err := r.Renderer.Clear(); err != nil {
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

	fmt.Printf("RenderLine: %d, %v", number, l)

	winWidth, _ := r.Window.GetSize()

	r.Renderer.SetDrawColor(r.Theme.Foreground.R, r.Theme.Foreground.G, r.Theme.Foreground.B, r.Theme.Foreground.A)

	rect := &sdl.Rect{
		X: 0,
		Y: int32(number * uint(r.Theme.LineHeight)),
		W: winWidth,
		H: int32(r.Theme.LineHeight),
	}

	if err := r.Renderer.FillRect(rect); err != nil {
		panic(err)
	}

	gfx.StringRGBA(r.Renderer, int32(r.Theme.LineHeight), int32((r.Theme.LineHeight*uint8(number))+(r.Theme.LineHeight/2)), "GFX Demo", r.Theme.Background.R, r.Theme.Background.G, r.Theme.Background.B, r.Theme.Background.A)
}
