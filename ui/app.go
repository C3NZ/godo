package ui

import (
    "log"
    tui "github.com/gizak/termui"
)

// Build the UI application and initaliaze all UI components
func BuildApp() {
    // Initialize the ui
    if err := tui.Init(); err != nil {
        log.Fatalf("The UI could not be initialized due to: ", err) 
    }
    paragraph := CreateHeader()
    defer tui.Close()
    tui.Render(paragraph) 

}
