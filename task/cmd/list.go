package cmd

import (
	"fmt"

	"github.com/noornee/golang_practice_repo/task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the tasks added by user",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, _ := db.ListTasks()
		for i, v := range tasks {
			//fmt.Println(i, v)
			fmt.Printf("%d: %s\n", i+1, v.Value)
		}
	},
}
