package task

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zrross11/golang-task-manager/pkg/db"
)

// Task represents a single task
type Task struct {
	ID   int
	Name string
}

// AddTask adds a new task to the task manager
func AddTask(name string) error {
	_, err := db.DB.Exec("INSERT INTO tasks (name) VALUES (?)", name)

	// then print all tasks before returning err
	ListTasks()
	return err
}

// ListTasks lists all the tasks in the task manager
func ListTasks() {
	rows, err := db.DB.Query("SELECT id, name FROM tasks")
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
func RemoveTask(name string) error {
	_, err := db.DB.Exec("DELETE FROM tasks WHERE name = ?", name)

	// then print all tasks before returning err
	ListTasks()
	return err
}

// EditTask edits a task in the task manager based on ID and a new name
func EditTask(id int, name string) error {
	_, err := db.DB.Exec("UPDATE tasks SET name = ? WHERE id = ?", name, id)

	// then print all tasks before returning err
	ListTasks()
	return err
}