package ui

import (
    "log"
   
    "godo/ui/header"

    tui "github.com/gizak/termui"
    widgets "github.com/gizak/termui/widgets"
)

// Build the UI application and initaliaze all UI components
func BuildApp() {
    // Initialize the ui
    if err := tui.Init(); err != nil {
        log.Fatalf("The UI could not be initialized due to: ", err) 
    }

    defer tui.close()
    

}
