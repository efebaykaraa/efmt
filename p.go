package efmt

import (
	"fmt"
	"strings"
)

func (p *Printer) Println(a ...any) {
	fmt.Println(p.prefix + fmt.Sprint(a...))
}

func (p *Printer) Printf(format string, a ...any) {
	format = strings.ReplaceAll(format, "\n", "\n"+p.prefix)
	if len(format) > 0 && format[len(format)-len(p.prefix):] == p.prefix {
		format = format[:len(format)-len(p.prefix)]
	}
	fmt.Printf(p.prefix+format, a...)
}

func (p *Printer) Print(a ...any) {
	fmt.Print(p.prefix, fmt.Sprint(a...))
}