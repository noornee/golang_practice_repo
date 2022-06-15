package main

import (
	"fmt"

	"github.com/noornee/golang_practice_repo/task/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
