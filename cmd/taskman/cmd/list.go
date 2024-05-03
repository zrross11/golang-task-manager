package cmd

import (
    "github.com/spf13/cobra"
	"github.com/zrross11/golang-task-manager/pkg/task"
)

var ListCmd = &cobra.Command{
    Use:   "list",
    Aliases: []string{"l"},
    Short:  "Lists tasks in the manager",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		task.ListTasks()
    },
}
