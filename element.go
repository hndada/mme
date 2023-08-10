package mme

type Element struct {
	*Node
	Display int // block, inline
	Layout  int // normal, flex, grid

	// Todo: move fields of Box to Element
	Box

	// Overflow: fixed to `auto``
	// Overflow[2]
	IsScrollAlwaysVisible bool
}
