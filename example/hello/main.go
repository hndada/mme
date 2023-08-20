package main

import (
	. "github.com/hndada/mme"
	"github.com/hndada/mme/tag"
)

func main() {
	app()
}

func app() *Element {
	a := tag.Anchor()
	c := 0
	a.OnClick = func() { c++ }
	a.Text.Data = "number: " + c
	return &a
}

func _() {

}
