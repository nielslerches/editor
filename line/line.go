package line

import (
	"math"
	"sort"
)

type Line struct {
	Content string
	Cursors []*Cursor
}

type Cursor struct {
	Offset uint
}

func NewLine() (l *Line) {
	l = &Line{
		Content: "",
		Cursors: make([]*Cursor, 0),
	}

	return l
}

func (l *Line) Split(index int) (b *Line) {
	b = &Line{Content: l.Content[index+1:]}
	l.Content = l.Content[:index]
	return b
}

func (l *Line) GetCursorIndex(c *Cursor) int {
	if c == nil {
		return -1
	}

	index := sort.Search(len(l.Cursors), func(i int) bool {
		return l.Cursors[i] == c
	})

	if index < len(l.Cursors) && l.Cursors[index] == c {
		return index
	}

	return -1
}

func (l *Line) MoveCursor(c *Cursor, offset int) {
	oldOffset := c.Offset
	newOffset := uint(math.Min(math.Max(float64(int(oldOffset)+offset), 0), float64(len(l.Content))))

	c.Offset = newOffset
}
