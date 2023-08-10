package mme

// https://developer.mozilla.org/en-US/docs/Learn/CSS/Styling_text/Fundamentals
type Text struct {
	Data  string
	Font  Font
	Color Color
	Size  Value

	WritingMode Direction
	Transform   TextTransform

	Underline     bool
	Strikethrough bool
	Align         Align
	LineHeight    Value
	LetterSpacing Value
	WordSpacing   Value
	WhiteSpace    WhiteSpace
	Shadows       []Shadow
}

type WhiteSpace int

const (
	WhiteSpaceNormal WhiteSpace = iota
	WhiteSpacePre
	WhiteSpaceNowrap
	WhiteSpacePreWrap
	WhiteSpacePreLine
)

type TextTransform int

const (
	TextTransformNone TextTransform = iota
	TextTransformCapitalize
	TextTransformUppercase
	TextTransformLowercase
	TextTransformFullWidth
)

type Direction int

const (
	LeftToRight Direction = iota
	RightToLeft
	TopToBottom
	BottomToTop
)
