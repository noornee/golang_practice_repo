package main

import (
	"fmt"

	"github.com/noornee/golang_practice_repo/task/cmd"
	"github.com/noornee/golang_practice_repo/task/db"
)

func main() {
	errHandler(db.InitDB("task.db"))
	errHandler(cmd.RootCmd.Execute())
}

func errHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
