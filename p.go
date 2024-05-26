package efmt

import (
	"fmt"
	"strings"
)

func (p *Printer) Println(a ...any) {
	fmt.Println(p.prefix + fmt.Sprint(a...))
}

func (p *Printer) Printf(format string, a ...any) {
	format = strings.Replace(format, "\n", "\n"+p.prefix, -1)
	if len(format) > 0 && format[len(format)-len(p.prefix):] == p.prefix {
		format = format[:len(format)-len(p.prefix)]
	}
	fmt.Printf(p.prefix+format, a...)
}

func (p *Printer) Print(s string) {
	fmt.Print(p.prefix + s)
}