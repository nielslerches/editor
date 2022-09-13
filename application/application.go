package application

import (
	"fmt"

	"github.com/nielslerches/editor/fonts"
	"github.com/nielslerches/editor/renderer"
	"github.com/nielslerches/editor/screen"
	"github.com/nielslerches/editor/theme"
	"github.com/veandco/go-sdl2/sdl"
)

type Application struct {
	Screen      *screen.Screen
	Renderer    *renderer.Renderer
	Theme       *theme.Theme
	FontManager *fonts.FontManager
}

func NewApplication() (a *Application) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	var f *fonts.Font

	fm := fonts.NewFontManager()
	fs := fm.GetFonts().FilterByStyle("monospace")

	if len(fs.Result) > 0 {
		f = fs.Result[0]
	}

	fmt.Printf("Font: %v\n", *f)

	t := &theme.Theme{
		Background: theme.NewColor(63, 63, 63, 255),
		Foreground: theme.NewColor(127, 127, 127, 255),
		LineHeight: 20,
		FontSize:   16,
		Font:       f,
		Text:       theme.NewColor(7, 7, 7, 255),
	}

	r := renderer.NewRenderer(t)

	a = &Application{
		Screen:      screen.NewScreen(),
		Renderer:    r,
		Theme:       t,
		FontManager: fm,
	}

	a.Render()

	return a
}

func (a *Application) Destroy() {
	a.Renderer.Destroy()
	sdl.Quit()
}

func (a *Application) Run() {
	for event := sdl.WaitEvent(); event != nil; event = sdl.WaitEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			goto quit
		case *sdl.KeyboardEvent:
			if t.State != sdl.PRESSED {
				continue
			}

			keyCode := t.Keysym.Sym
			fmt.Printf("%v", keyCode)

			switch keyCode {
			case sdl.K_ESCAPE:
				goto quit
			case sdl.K_BACKSPACE:
				if a.Screen.SelectedEditor != nil {
					l := a.Screen.SelectedEditor.GetLineByIndex(0)
					if l != nil {
						if len(l.Content) > 0 {
							l.Content = l.Content[:len(l.Content)-1]
						} else {
							l.Content = ""
						}
					}
				}
			}
		case *sdl.WindowEvent:
			switch int(t.Event) {
			case sdl.WINDOWEVENT_RESIZED:
				a.OnWindowResize()
			}
		case *sdl.TextInputEvent:
			fmt.Println("TextInputEvent")
			if a.Screen.SelectedEditor != nil {
				l := a.Screen.SelectedEditor.GetLineByIndex(0)
				if l != nil {
					l.Content = l.Content + t.GetText()
				}
			}
		}
		a.Render()
	}

quit:
	a.Destroy()
}

func (a *Application) Render() {
	a.Renderer.Render(a.Screen)
}

func (a *Application) OnWindowResize() {
	w, h := a.Renderer.Window.GetSize()
	fmt.Printf("OnWindowResize: %d, %d", w, h)
	a.Renderer.Resize(w, h)
}
