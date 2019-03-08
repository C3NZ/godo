package ui

import (
	"strconv"
	"strings"

	"github.com/C3NZ/godo/db"

	clui "github.com/VladimirMarkelov/clui"
	term "github.com/nsf/termbox-go"
)

var todoList *clui.ListBox
var currentTask string
var taskMsg string = "Enter task:"

// CreateTodoWindow creates a todo window
func CreateTodoWindow() {
	todoWindow := clui.AddWindow(0, 0, 50, 50, "godo list")
	todoWindow.SetPack(1)
	todoWindow.SetPaddings(4, 3)
	todoWindow.SetGaps(0, 2)

	createTodoList(todoWindow)
	createTodoEdit(todoWindow)
}

// CreateTodoList creates the Listbox for displaying the todos
func createTodoList(view *clui.Window) {
	todoList = clui.CreateListBox(view, 30, 40, 100)
	todoList.OnSelectItem(func(e clui.Event) {
		item := todoList.FindItem(e.Msg, true)
		if item != -1 {
			todoList.RemoveItem(item)
		}
	})

	tasks, err := db.AllTasks()
	if err != nil {
		panic(err)
	}
	for _, task := range tasks {
		num := strconv.Itoa(task.Key)
		todoList.AddItem(num + "." + task.Value)
	}
}

// CreateTodoEdit creates the Edit field for the adding todos
func createTodoEdit(view *clui.Window) {
	todoEdit := clui.CreateEditField(view, 30, taskMsg, 100)

	// Register the todo edit on change
	todoEdit.OnChange(func(e clui.Event) {
		currentTask = e.Msg
	})

	// Register the todo edit on key press
	todoEdit.OnKeyPress(func(key term.Key, testRune rune) bool {
		if key == term.KeyEnter {
			currentTask = strings.Split(currentTask, ":")[1]
			AddItem(currentTask)
			todoEdit.Clear()
			todoEdit.SetTitle(taskMsg)
			return true
		}

		return false
	})
}

// AddItem adds an item to the task list
func AddItem(item string) {
	id, err := db.CreateTask(item)
	if err != nil {
		panic(err)
	}

	todoList.AddItem(strconv.Itoa(id) + "." + item)
}

// DeleteItem deletes an item from the db
func DeleteItem(id int) {
	err := db.DeleteTask(id)
	if err != nil {
		panic(err)
	}
}
