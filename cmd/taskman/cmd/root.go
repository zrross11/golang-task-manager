package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zrross11/golang-task-manager/pkg/db"
)

// NewRootCmd creates a new root command for taskman. It is called once in the
// main function.
func NewRootCmd() *cobra.Command {

	rootCmd := &cobra.Command{
		Use:     "taskman",
		Short:   "Task Manager CLI",
		Long:    "Task Manager CLI.\n\nTaskman is a lightweight CLI to manage tasks.",
		Version: "0.0.1",
	}
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	initRootCmd(rootCmd)

	return rootCmd
}

func initRootCmd(rootCmd *cobra.Command) {
	// Initialize the database
	var err error
	db.DB, err = db.InitializeDB("../../../tasks.db")
	if err != nil {
		fmt.Println("Error inserting task:", err)
		return
	}

	// Add subcommands
	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(RemoveCmd)
	rootCmd.AddCommand(EditCmd)
}

// Execute runs the root command. It is called once in the main function.
func Execute(rootCmd *cobra.Command) error {
	// Use the db object for further operations
	return rootCmd.Execute()
}
