package github.com/efexplose/efmt

import (
	"fmt"
	"strings"
)

type Printer struct {
	prefix string
}

func NewPrinter(prefix string) *Printer {
	return &Printer{
		prefix: prefix,
	}
}