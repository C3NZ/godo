package ui

import (
	widgets "github.com/gizak/termui/widgets"
)

// CreateHeader creates the header for godo
func CreateHeader() *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Text = "godo"
	return p
}
