package mme

import "fmt"

type Color string

var (
	Blue   Color = "#0000FF"
	Gray94 Color = "#F0F0F0"
)

func RGB(r, g, b uint8) string {
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}
func RGBA(r, g, b, a uint8) string {
	return fmt.Sprintf("#%02x%02x%02x%02x", r, g, b, a)
}
