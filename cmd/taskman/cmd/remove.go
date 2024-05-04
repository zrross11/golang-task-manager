package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zrross11/golang-task-manager/pkg/task"
)

var RemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"r"},
	Short:   "Remove a task from the manager",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a task to remove:")
			reader := bufio.NewReader(os.Stdin)
			taskName, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading user input")
				return
			}
			taskName = taskName[:len(taskName)-1]
			task.RemoveTask(taskName)
			return
		} else {
			task.RemoveTask(args[0])
		}
	},
}
