package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager...more like a todo CLI task manager",
}

func init() {
	RootCmd.AddCommand(addCmd, listCmd, doCmd)
}
