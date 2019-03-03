package ui

import (
    widgets "github.com/gizak/termui/widgets"
)

func CreateHeader() *widgets.Paragraph {
    p := widgets.NewParagraph()
    p.Text = "Hello connor"    
    p.SetRect(40, 0, 70, 20)

    return p
}
