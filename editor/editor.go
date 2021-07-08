package editor

import (
	"sort"

	"github.com/nielslerches/editor/line"
)

type Editor struct {
	Lines []*line.Line
}

func NewEditor() (e *Editor) {
	l := line.NewLine()
	ls := []*line.Line{l}

	e = &Editor{
		Lines: ls,
	}

	return e
}

func (e *Editor) SplitLine(l *line.Line, index int) {
	newLine := l.Split(index)
	e.Lines = append(append(e.Lines[:index], newLine), e.Lines[:index+1]...)
}

func (e *Editor) GetLineIndex(l *line.Line) int {
	if e == nil {
		return -1
	}

	indexOfLine := sort.Search(len(e.Lines), func(i int) bool {
		return e.Lines[i] == l
	})

	if indexOfLine < len(e.Lines) && e.Lines[indexOfLine] == l {
		return indexOfLine
	}

	return -1
}

func (e *Editor) GetLineByIndex(indexOfLine int) *line.Line {
	if indexOfLine < len(e.Lines) && indexOfLine > -1 {
		return e.Lines[indexOfLine]
	}

	return nil
}
