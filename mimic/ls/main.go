package main

import (
	"io/ioutil"
	"os"

	"github.com/noornee/golang_practice_repo/mimic/ls/custom"
)

// TODO: Implement recursion for nested folders
func main() {
	arg := os.Args[1:]
	files, err := ioutil.ReadDir(".")

	if len(arg) == 0 {
		custom.CustomError(err)
		for _, value := range files {
			custom.CustomValue(value.Name())
		}

	} else {

		files, err := ioutil.ReadDir(arg[0])
		custom.CustomError(err)

		for _, value := range files {
			custom.CustomValue(value.Name())
		}
	}
}
