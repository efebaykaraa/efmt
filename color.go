package efmt

import "fmt"

type Color struct {
	R, G, B int
}

var (
	Red   = Color{255, 0, 0}
	Green = Color{0, 255, 0}
	Blue  = Color{0, 0, 255}
	White = Color{255, 255, 255}
	Black = Color{0, 0, 0}
)

func NewColor(r, g, b int) Color {
	return Color{r, g, b}
}

func (color Color) Render(s string) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", color.R, color.G, color.B, s)
}

func (color Color) String() string {
	return fmt.Sprintf(color.Render("Color(%d, %d, %d)"), color.R, color.G, color.B)
}