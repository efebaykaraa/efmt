package efmt

import (
	"testing"
)

func TestPrinter(t *testing.T) {
	p := NewPrinter("[Test] ")
	p.Print("Hello World\n")
	p.Println("Hello World")
	p.Printf("Hello %s\n", "World")
}

func TestColor(t *testing.T) {
	p := NewPrinter("[Test] ")
	p.Printc(Red, "Red Text\n")
	p.Printc(Blue, "Blue: Text\n")

	p.Printlnc(Red, "Red Text")
	p.Printlnc(Blue, "Blue: Text")

	p.Printfc(Red, "Red %s\n", "Text")
	p.Printfc(Blue, "Blue: %s\n", "Text")
}