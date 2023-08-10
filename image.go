package mme

type Image interface{}

// Includes an image or gradients.
type Picture struct {
	Source      string
	Description string
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
	Stop  Value
}
