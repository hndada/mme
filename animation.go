package mme

type Animation struct {
	Name         string
	Duration     Value
	Easing       func()
	Delay        Value
	MaxIteration int
	Direction    Direction
	FillMode     FillMode
	IsPaused     bool
}

type FillMode int

const (
	FillModeNone FillMode = iota
	FillModeForwards
	FillModeBackwards
	FillModeBoth
)
