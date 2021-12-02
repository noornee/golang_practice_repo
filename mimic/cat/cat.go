package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		
		val := fmt.Sprintf("%v", args[i])
		file, err := os.ReadFile(val)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%v\n", string(file))

	}

}
