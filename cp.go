package github.com/efexplose/efmt

func (p *Printer) Printlnc(color Color, a ...any) {
	p.Println(colorizeToPrint(color, strings.Split(fmt.Sprint(a...), "\n")))
}

func (p *Printer) Printfc(color Color, format string, a ...any) {
	p.Print(colorizeToPrint(color, strings.Split(fmt.Sprintf(format, a...), "\n")))
}

func colorizeToPrint(color Color, lines []string) string {
	var result string
	for _, line := range lines {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			result += color(parts[0]) + ": " + parts[1] + "\n"
		} else {
			result += line + "\n"
		}
	}
	return result
}