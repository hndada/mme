package mme

import "fmt"

type Value string

type UnitType int

const (
	// Absolute length units
	Pixel UnitType = iota
	Inch
	Pica
	Point
	Centimeter
	Millimeter
	QuarterMillimeter

	// Relative length units
	Percent
	// relative to font-size of parent. Stands for the width of the letter M.
	Em
	RootEm
	ViewportWidth
	ViewportHeight
	ViewportMin
	ViewportMax

	Degree
	Radian
	Gradian
)

var unitSuffixes = map[UnitType]string{
	Pixel:             "px",
	Inch:              "in",
	Pica:              "pc",
	Point:             "pt",
	Centimeter:        "cm",
	Millimeter:        "mm",
	QuarterMillimeter: "Q",

	Percent:        "%",
	Em:             "em",
	RootEm:         "rem",
	ViewportWidth:  "vw",
	ViewportHeight: "vh",
	ViewportMin:    "vmin",
	ViewportMax:    "vmax",

	Degree:  "deg",
	Radian:  "rad",
	Gradian: "grad",
}

// both "35px" and Unit(35, Pixel) are valid.
// Unit(35, Pixel) returns "35px"
func Unit(v float64, t UnitType) string { return fmt.Sprintf("%f%s", v, unitSuffixes[t]) }

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
