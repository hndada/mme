package mme

// For sides.
const (
	Top = iota
	Right
	Bottom
	Left
)

// For corners.
const (
	TopLeft = iota
	TopRight
	BottomRight
	BottomLeft
)

// type Size struct{ W, H Length }
// type Position struct{ X, Y Length }
type Size struct{ W, H Length }
type XY struct{ X, Y Length }

// For Line.Style
const (
	Dashed = iota
)

type Line struct {
	Thickness Length // width
	Style     string
	Color     Color
}

// Todo: box-sizing
type Box struct {
	Size
	Min Size
	Max Size
	XY

	// Boundary
	Margin  [4]Value
	Border  [4]Line
	Radius  [4]XY
	Padding [4]Value

	Background
}

type Background struct {
	Images []Image
}

type Image interface{}

// Includes an image or gradients.
type Picture struct {
	Source string
	Size
	XY
	Repeat [2]bool
}

const (
	LinearGradient = iota
	RadialGradient
)

type Gradient struct {
	Type  int // LinearGradient, RadialGradient
	Stops []GradientStop
}

type GradientStop struct {
	Color Color
	Stop  Length
}
