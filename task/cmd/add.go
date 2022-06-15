package cmd

import (
	"fmt"
	"strings"

	"github.com/noornee/golang_practice_repo/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds task to task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("there was an error creating task", err)
			return
		}
		fmt.Printf("successfully added \"%s\" to task list\n", task)
	},
}
