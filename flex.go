package mme

type Axis int

const (
	AxisX Axis = iota
	AxisY
)

// https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_flexible_box_layout/Basic_concepts_of_flexbox
// The items do not stretch on the main dimension, but can shrink.
// The items will stretch to fill the size of the cross axis.

// Each element's size substitutes Basis, Grow, Shrink.
type FlexibleBox struct {
	GrowingAxis         Axis
	GrowingAxisReversed bool
	Wrap                bool
	WrapReversed        bool
	Stretched           bool
	Align               Align
	Justify             Justify
}

type Align int

const (
	AlignStart Align = iota
	AlignCenter
	AlignEnd
)

type Justify int

const (
	JustifyStart Justify = iota
	JustifyCenter
	JustifyEnd

	JustifySpaceBetween
	JustifySpaceAround
	JustifySpaceEvenly
)
