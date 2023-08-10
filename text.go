package mme

type Text struct {
	Data        string
	WritingMode Direction
}

type Direction int

const (
	LeftToRight Direction = iota
	RightToLeft
	TopToBottom
	BottomToTop
)
