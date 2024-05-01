package task

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Task represents a single task
type Task struct {
	ID   int
	Name string
}

// TaskManager represents the task manager
type TaskManager struct {
	db *sql.DB
}

// NewTaskManager creates a new task manager
func NewTaskManager(db *sql.DB) *TaskManager {
	return &TaskManager{
		db: db,
	}
}

// AddTask adds a new task to the task manager
func (tm *TaskManager) AddTask(name string) error {
	_, err := tm.db.Exec("INSERT INTO tasks (name) VALUES (?)", name)

	// then print all tasks before returning err
	tm.ListTasks()
	return err
}

// ListTasks lists all the tasks in the task manager
func (tm *TaskManager) ListTasks() {
	rows, err := tm.db.Query("SELECT id, name FROM tasks")
	if err != nil {
		fmt.Println("Error querying tasks:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Tasks:")

	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Name)
		if err != nil {
			fmt.Println("Error scanning tasks:", err)
			return
		}
		fmt.Printf("Id: %d. Task: %s\n", t.ID, t.Name)
	}
	fmt.Println()
}

// RemoveTask removes a task from the task manager
func (tm *TaskManager) RemoveTask(name string) error {
	_, err := tm.db.Exec("DELETE FROM tasks WHERE name = ?", name)

	// then print all tasks before returning err
	tm.ListTasks()
	return err
}

// RunCLI runs the CLI task manager
func RunCLI(db *sql.DB) {
	fmt.Println("Starting Task Manager:")
	TaskManager := NewTaskManager(db)

	for {

		fmt.Println("1. Add a task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Remove a task")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scanln(&choice)

		if choice == 4 {
			fmt.Println("Exiting...")
			break
		} else if choice == 2 {
			TaskManager.ListTasks()
		} else if choice == 1 {

			// Read input of task name from console
			fmt.Println("Enter task to add:")
			var name string
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				name = scanner.Text()
			}
			if err := scanner.Err(); err != nil {
				fmt.Println("Error scanning name:", err)
				return
			}
			//end scanning

			TaskManager.AddTask(name)
		} else if choice == 3 {
			// Read input of task name from console
			fmt.Println("Enter task to remove:")
			var name string
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				name = scanner.Text()
			}
			if err := scanner.Err(); err != nil {
				fmt.Println("Error scanning name:", err)
				return
			}
			//end scanning

			TaskManager.RemoveTask(name)
		} else {
			fmt.Println("Invalid choice")
		}
	}
}
