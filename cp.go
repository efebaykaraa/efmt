package efmt

import (
	"fmt"
)

func (p *Printer) Printc(color Color, s string) {
	p.Print(colorizeToPrint(color, s))
}

func (p *Printer) Printlnc(color Color, a ...any) {
	p.Println(colorizeToPrint(color, fmt.Sprint(a...)))
}

func (p *Printer) Printfc(color Color, format string, a ...any) {
	p.Print(colorizeToPrint(color, fmt.Sprintf(format, a...)))
}