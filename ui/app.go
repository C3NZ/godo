package ui

import (
	clui "github.com/VladimirMarkelov/clui"
	"github.com/luisvinicius167/godux"
)

// Store is the ui application state store for managing the application state
var Store *godux.Store = godux.NewStore()

// BuildApp builds the UI application and initaliaze all UI components
func BuildApp() {
	// Initialize the ui
	clui.InitLibrary()

	CreateTodoWindow()
	defer clui.DeinitLibrary()
	clui.MainLoop()
}
