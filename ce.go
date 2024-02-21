package efmt

import (
	"fmt"
)

func (p *Printer) Errorfc(color Color, format string, a ...any) error {
	return fmt.Errorf(colorizeToPrint(color, fmt.Sprintf(format, a...)))
}

func (p *Printer) Errorlnc(color Color, a ...any) error {
	return p.Errorln(colorizeToPrint(color, fmt.Sprint(a...)))
}