package tests

import (
	"github.com/C3NZ/godo/db"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

// Test our database
func TestDb(t *testing.T) {
	// Obtain the homedir path for the user currently running this program
	home, homeDirErr := homedir.Dir()
	if homeDirErr != nil {
		t.Errorf("The home directory could not be accessed: %s", homeDirErr)
	}

	// Create the test DB path
	dbPath := filepath.Join(home, "testDB.db")

	dbInitErr := db.Init(dbPath)
	if dbInitErr != nil {
		t.Errorf("Couldnt initialize the database: %s", dbInitErr)
	}
}

func TestAddItems(t *testing.T) {
	var tasks = []string{
		"This is a task",
		"This is anotheer task",
		"This is the third task",
	}

	// Iterate over the pseudo tasks
	for iter, task := range tasks {
		taskId, taskErr := db.CreateTask(task)
		// Check if an error occurred
		if taskErr != nil {
			t.Errorf("Couldnt Create task \"%d \": %s", iter, taskErr)
		}

		// Check if task id matches up
		if taskId != iter {
			t.Errorf("The task was not created with the expected ID: %d, instead got: %d", iter, taskId)
		}

	}
}
