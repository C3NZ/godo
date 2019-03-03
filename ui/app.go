package ui

import (
    "log"
   
    gui "github.com/C3NZ/godo/ui"

    tui "github.com/gizak/termui"
)

// Build the UI application and initaliaze all UI components
func BuildApp() {
    // Initialize the ui
    if err := tui.Init(); err != nil {
        log.Fatalf("The UI could not be initialized due to: ", err) 
    }
    paragraph := ui.CreateHeader()
    defer tui.close()
    tui.Render(paragraph) 

}
