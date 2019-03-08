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
var taskMsg = "Enter task:"
var taskMap = make(map[int]int)

// CreateTodoWindow creates a todo window
func CreateTodoWindow() {
	todoWindow := clui.AddWindow(0, 0, 50, 30, "godo list")
	todoWindow.SetPack(1)
	todoWindow.SetPaddings(4, 3)
	todoWindow.SetGaps(0, 2)
	todoWindow.SetMaximized(true)
	todoWindow.SetSizable(false)

	createTodoList(todoWindow)
	createTodoEdit(todoWindow)
}

func updateTasks() {
	tasks, err := db.AllTasks()
	if err != nil {
		panic(err)
	}
	for id, task := range tasks {
		taskMap[id+1] = task.Key
		num := strconv.Itoa(id + 1)
		todoList.AddItem(num + "." + task.Value)
	}
}

// CreateTodoList creates the Listbox for displaying the todos
func createTodoList(view *clui.Window) {
	todoList = clui.CreateListBox(view, 30, 20, 50)
	todoList.OnSelectItem(func(e clui.Event) {
		item := todoList.FindItem(e.Msg, true)
		if item != -1 {
			strID := string(strings.Split(e.Msg, " ")[0][0])
			intID, err := strconv.Atoi(strID)
			if err != nil {
				panic(err)
			}
			todoList.RemoveItem(item)
			todoList.Clear()
			DeleteItem(taskMap[intID])
			updateTasks()

		}
	})
	updateTasks()
}

// CreateTodoEdit creates the Edit field for the adding todos
func createTodoEdit(view *clui.Window) {
	todoEdit := clui.CreateEditField(view, 30, taskMsg, 50)

	// Register the todo edit on change
	todoEdit.OnChange(func(e clui.Event) {
		currentTask = e.Msg
	})

	// Register the todo edit on key press
	todoEdit.OnKeyPress(func(key term.Key, testRune rune) bool {
		if key == term.KeyEnter {
			currentTask = strings.Split(currentTask, ":")[1]
			AddItem(currentTask)
			todoList.Clear()
			updateTasks()
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
