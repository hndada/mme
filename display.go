package mme

type DisplayType int

const (
	// Block: newline, w, h, padding, margin, border are respected
	// fills 100% when width is not provided
	DisplayBlock DisplayType = iota

	// Inline: left, right are respected (padding, margin, border)
	// top, bottom are partially respected
	DisplayInline
)

type LayoutType int

const (
	// normal flow
	LayoutNormal LayoutType = iota

	// Diplay type: block, Layout model: flexbox
	LayoutFlex

	LayoutGrid
)
