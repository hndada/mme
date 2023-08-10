package mme

type Type int

func (t Type) String() string {
	switch t {
	case Pixel:
		return "px"
	}
	return ""
}

const (
	Pixel Type = iota
)

// Centimeter(10)
// 10*Centimeter
// "10cm"
type Length struct {
	Type
}

// "The only value that you will commonly use is pixels."
// https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Values_and_units
// type Pixel float64

// const (
// 	Inch  Pixel = 96
// 	Pica  Pixel = Inch / 6
// 	Point Pixel = Pica / 12

// 	Centimeter        Pixel = 37.8
// 	Millimeter        Pixel = Centimeter / 10
// 	QuarterMillimeter Pixel = Millimeter / 4
// )
