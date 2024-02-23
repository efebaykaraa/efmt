package efmt

import (
	"fmt"
	"strings"
)

type Color struct {
	R, G, B int
}

var (
	// basic
	Red   = Color{255, 0, 0}
	Green = Color{0, 255, 0}
	Blue  = Color{0, 0, 255}
	Yellow = Color{255, 255, 0}
	Cyan   = Color{0, 255, 255}
	Magenta = Color{255, 0, 255}
	White = Color{255, 255, 255}
	Black = Color{0, 0, 0}

	// gray
	Gray = Color{128, 128, 128}
	LightGray = Color{192, 192, 192}
	DarkGray = Color{64, 64, 64}

	// light
	LightRed = Color{255, 128, 128}
	LightGreen = Color{128, 255, 128}
	LightBlue = Color{128, 128, 255}
	LightYellow = Color{255, 255, 128}
	LightCyan = Color{128, 255, 255}
	LightMagenta = Color{255, 128, 255}

	// dark
	DarkRed = Color{128, 0, 0}
	DarkGreen = Color{0, 128, 0}
	DarkBlue = Color{0, 0, 128}
	DarkYellow = Color{128, 128, 0}
	DarkCyan = Color{0, 128, 128}
	DarkMagenta = Color{128, 0, 128}
)

func NewColor(r, g, b int) Color {
	return Color{r, g, b}
}

func HexColor(hex int) Color {
	return Color{
		R: (hex >> 16) & 0xFF,
		G: (hex >> 8) & 0xFF,
		B: hex & 0xFF,
	}
}

func (color Color) Render(s string) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", color.R, color.G, color.B, s)
}

func colorizeToPrint(color Color, text string) string {
	var result string
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			result += color.Render(parts[0] + ": ") + parts[1]
		} else {
			result += color.Render(line)
		}
	}
	return result
}

func (color Color) String() string {
	return color.Render(fmt.Sprintf("{R: %d, G: %d, B: %d}", color.R, color.G, color.B))
}
