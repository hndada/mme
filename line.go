package mme

// For Line.Style
const (
	Dashed = iota
)

type Line struct {
	Thickness Value // width
	Style     string
	Color     Color
}
