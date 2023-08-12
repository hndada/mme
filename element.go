package mme

// Todo: box-sizing
// Todo: overflow
type Element struct {
	Tag string // for semantic tags

	*Node
	Display DisplayType
	Layout  LayoutType

	Size
	Min      Size
	Max      Size
	Position XY

	Border     Border
	Background Background
	Text       Text
	Shadow     Shadow

	IsScrollAlwaysVisible bool

	// Interaction
	Cursor     CursorType
	CaretColor Color
}

type CursorType int

const (
	CursorDefault CursorType = iota
	CursorPointer
)
