package cmd

import (
    "github.com/spf13/cobra"
	"github.com/zrross11/golang-task-manager/pkg/task"
)

var AddCmd = &cobra.Command{
    Use:   "add",
    Aliases: []string{"a"},
    Short:  "Add a task to the manager",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		task.AddTask(args[0])
    },
}
