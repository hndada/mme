package mme

type FontStyle int

const (
	FontStyleNormal FontStyle = iota
	FontStyleItalic
	FontStyleOblique
)

type FontWeight int

const (
	FontWeightNormal FontWeight = iota
	FontWeightBold
	FontWeightBolder
	FontWeightLighter
	// 100~900
)

type Font struct {
	Family []string
	Style  FontStyle
	Weight FontWeight
	Size
}

func NewFont(family string) Font {
	return Font{
		Family: []string{family},
	}
}
