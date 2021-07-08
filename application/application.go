package application

import (
	"fmt"

	"github.com/nielslerches/editor/renderer"
	"github.com/nielslerches/editor/screen"
	"github.com/nielslerches/editor/theme"
	"github.com/veandco/go-sdl2/sdl"
)

type Application struct {
	Screen   *screen.Screen
	Renderer *renderer.Renderer
	Theme    *theme.Theme
}

func NewApplication() (a *Application) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	t := theme.DEFAULT

	r := renderer.NewRenderer(t)

	a = &Application{
		Screen:   screen.NewScreen(),
		Renderer: r,
		Theme:    t,
	}

	a.Render()

	return a
}

func (a *Application) Destroy() {
	a.Renderer.Destroy()
	sdl.Quit()
}

func (a *Application) Run() {
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if t.State != sdl.PRESSED {
					continue
				}

				keyCode := t.Keysym.Sym

				switch keyCode {
				case sdl.K_ESCAPE:
					running = false
				}
			case *sdl.WindowEvent:
				switch int(t.Event) {
				case sdl.WINDOWEVENT_RESIZED:
					a.OnWindowResize()
				}
			}

			a.Render()
		}
	}
}

func (a *Application) Render() {
	a.Renderer.Render(a.Screen)
}

func (a *Application) OnWindowResize() {
	w, h := a.Renderer.Window.GetSize()
	fmt.Printf("OnWindowResize: %d, %d", w, h)
	a.Renderer.Resize(w, h)
}
