package cmd

import (
	"fmt"
    "strings"
    "github.com/C3NZ/godo/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use: "add",
    Short: "Add a task to the to do list",
    Run: func(cmd *cobra.Command, args []string) {

        task := strings.Join(args, " ")
        taskId, err := db.CreateTask(task)
        if err != nil {
            fmt.Printf("A sudden error occurred: %s\n", err.Error())
        }
        fmt.Printf("Added \"%s\" to your task list with an id of: %d\n", task, taskId)
    },
}

// Go initializer function called to setup anything before execution
func init() {
    // Register the addCmD as a sub command
    RootCmd.AddCommand(addCmd)
}
