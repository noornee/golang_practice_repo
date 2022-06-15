package cmd

import (
	"fmt"
	"strconv"

	"github.com/noornee/golang_practice_repo/task/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks task as done and deletes from database",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("do called")
		//err := db.DeleteTask(args[0])
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("there was an error parsing arg", err)
			} else {
				ids = append(ids, id)
			}
		}
		//fmt.Println(ids)
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("something went wrong", err)
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("invalid task number")
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			fmt.Println(task.Key, err.Error())
			//if err != nil {
			//fmt.Println("error deleting task: ", err)
			//} else {
			//fmt.Printf("successfully deleted task with id of \"%d\"\n", task.Key)
			//}

		}
	},
}
