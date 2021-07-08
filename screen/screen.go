package screen

import (
	"sort"

	"github.com/nielslerches/editor/editor"
)

type Screen struct {
	Editors        []*editor.Editor
	SelectedEditor *editor.Editor
}

func NewScreen() (s *Screen) {
	e := editor.NewEditor()
	s = &Screen{
		Editors:        []*editor.Editor{e},
		SelectedEditor: e,
	}

	return s
}

func (s *Screen) SelectEditor(e *editor.Editor) {
	s.SelectedEditor = e
}

func (s *Screen) CloseEditor(e *editor.Editor) {
	if e == nil {
		return
	}

	indexOfEditor := s.GetEditorIndex(e)

	if indexOfEditor > -1 {
		if len(s.Editors) > 1 {
			s.SelectEditor(s.Editors[indexOfEditor-1])
		}

		// Remove element from array at index
		s.Editors = append(s.Editors[:indexOfEditor], s.Editors[indexOfEditor+1:]...)
	}

	if len(s.Editors) == 0 {
		s.SelectEditor(nil)
	}
}

func (s *Screen) GetEditorIndex(e *editor.Editor) int {
	if e == nil {
		return -1
	}

	indexOfEditor := sort.Search(len(s.Editors), func(i int) bool {
		return s.Editors[i] == e
	})

	if indexOfEditor < len(s.Editors) && s.Editors[indexOfEditor] == e {
		return indexOfEditor
	}

	return -1
}
