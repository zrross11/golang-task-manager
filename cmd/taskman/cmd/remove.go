package cmd

import (
    "github.com/spf13/cobra"
	"github.com/zrross11/golang-task-manager/pkg/task"
)

var RemoveCmd = &cobra.Command{
    Use:   "remove",
    Aliases: []string{"r"},
    Short:  "Remove a task from the manager",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		task.RemoveTask(args[0])
    },
}
