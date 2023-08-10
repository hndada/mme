package mme

type Size struct {
	Width  Value
	Height Value
}

type XY struct{ X, Y Value }

// For sides
const (
	SideTop = iota
	SideRight
	SideBottom
	SideLeft
)

func AllSides(v Value) [4]Value              { return [4]Value{v, v, v, v} }
func VerticalHorizontal(v, h Value) [4]Value { return [4]Value{v, h, v, h} }

// For corners
const (
	CornerTopLeft = iota
	CornerTopRight
	CornerBottomRight
	CornerBottomLeft
)

func AllCorners(v Value) [4]XY {
	xy := XY{v, v}
	return [4]XY{xy, xy, xy, xy}
}
