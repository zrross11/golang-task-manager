package main

import (
	"github.com/zrross11/golang-task-manager/internal/task"
)

func main() {

	db, err := task.InitializeDB("tasks.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	task.RunCLI(db)
}
