package ui

import (
	"log"

	tui "github.com/gizak/termui"
	"github.com/luisvinicius167/godux"
)

// Store is the ui application state store for
var Store *godux.Store = godux.NewStore()

// BuildApp builds the UI application and initaliaze all UI components
func BuildApp() {
	// Initialize the ui
	if err := tui.Init(); err != nil {
		log.Fatalln("The UI could not be initialized due to: ", err)
	}

	paragraph := CreateHeader()
	defer tui.Close()
	tui.Render(paragraph)

}
