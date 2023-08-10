package mme

type Color struct {
	R, G, B, A uint8
}

var (
	Black = Color{0, 0, 0, 255}
)

func RGB(r, g, b uint8) Color     { return Color{r, g, b, 255} }
func RGBA(r, g, b, a uint8) Color { return Color{r, g, b, a} }
