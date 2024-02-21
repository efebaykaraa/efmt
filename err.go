package efmt

func (p *Printer) Errorf(format string, a ...any) error {
	format = strings.ReplaceAll(format, "\n", "\n"+p.prefix)
	if len(format) > 0 && format[len(format)-len(p.prefix):] == p.prefix {
		format = format[:len(format)-len(p.prefix)]
	}
	return fmt.Errorf(format, a...)
}

func (p *Printer) Errorln(a ...any) error {
	return Errorln(p.prefix + fmt.Sprint(a...))
}