package cmd

import (
    "github.com/spf13/cobra"
	"github.com/zrross11/golang-task-manager/pkg/task"
	"strconv"
)

var EditCmd = &cobra.Command{
	Use:   "edit",
	Aliases: []string{"e"},
	Short:  "Edit a task in the manager",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskId, _ := strconv.Atoi(args[0])
		task.EditTask(taskId, args[1])
	},
}
