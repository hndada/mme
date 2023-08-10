package mme

const (
	// Block: newline, w, h, padding, margin, border are respected
	// fills 100% when width is not provided
	DisplayBlock = iota

	// Inline: left, right are respected (padding, margin, border)
	// top, bottom are partially respected
	DisplayInline
)

const (
	// normal flow
	LayoutNormal = iota

	// Diplay type: block, Layout model: flexbox
	LayoutFlex

	LayoutGrid
)
