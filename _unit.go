package mme

import "fmt"

// type Value string // CSS value
// type Length string // CSS length value
// type Degree string // CSS angle value

// Proposal: 35 + Pixel (= 35px)?
// This may occur overflow.

// Proposal: Length{35, Pixel} (= 35px)?
// Proposal: Length(35, Percent) (= 35px)?
// type Unit string

type Unit struct {
	Value float64
	Type  int
}
type Length = Unit
type Angle = Unit

func NewLength(v float64) Unit { return Unit{Type: UnitPixel, Value: v} }

func ParseUnit(s string) Unit {
	return Unit{}
}

const (
	// Absolute length units
	UnitPixel = iota
	UnitInch
	UnitPica
	UnitPoint
	UnitCentimeter
	UnitMillimeter
	UnitQuarterMillimeter

	// Relative length units
	UnitPercent
	UnitEm // relative to font-size of parent. Stands for the width of the letter M.
	UnitRootEm
	UnitViewportWidth
	UnitViewportHeight
	UnitViewportMin
	UnitViewportMax

	UnitDegree
	UnitRadian
	UnitGradian
)

func (u Unit) String() string {
	var f = fmt.Sprintf
	switch u.Type {
	// Absolute length units
	case UnitPixel:
		return f("%fpx", u.Value)
	case UnitInch:
		return f("%fin", u.Value)
	case UnitPica:
		return f("%fpc", u.Value)
	case UnitPoint:
		return f("%fpt", u.Value)
	case UnitCentimeter:
		return f("%fcm", u.Value)
	case UnitMillimeter:
		return f("%fmm", u.Value)
	case UnitQuarterMillimeter:
		return f("%fQ", u.Value)

	// Relative length units
	case UnitPercent:
		return f("%f%%", u.Value)
	case UnitEm:
		return f("%fem", u.Value)
	case UnitRootEm:
		return f("%frem", u.Value)
	case UnitViewportWidth:
		return f("%fvw", u.Value)
	case UnitViewportHeight:
		return f("%fvh", u.Value)
	case UnitViewportMin:
		return f("%fvmin", u.Value)
	case UnitViewportMax:
		return f("%fvmax", u.Value)

	case UnitDegree:
		return f("%fdeg", u.Value)
	case UnitRadian:
		return f("%frad", u.Value)
	case UnitGradian:
		return f("%fgrad", u.Value)
	}
	return ""
}
