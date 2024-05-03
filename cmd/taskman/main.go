package main

import (
	"github.com/zrross11/golang-task-manager/cmd/taskman/cmd"
	"os"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}