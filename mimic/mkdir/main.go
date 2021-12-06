package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	for i, _ := range args {
		err := os.Mkdir(args[i], 0755)

		if err != nil {
			fmt.Printf("cannot create directory '%v': File exists\n", args[i])
		}
	}

}
