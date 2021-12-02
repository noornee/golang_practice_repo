package custom

import (
	"fmt"
	"strings"
)

func CustomError(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func CustomValue(value string) {
	// checks if there are files with whitespaces in the name e.g. 	hello world.txt prints ==>  'hello world.txt'
	// for _, value := range files {
	// 	handlers.CustomValue(value.Name())
	// }
	if strings.Contains(value, " ") {
		fmt.Printf("'%v' ", value)
	} else {
		fmt.Printf("%v ", value)
	}
}




/*
func main() {

	var Blue = "\033[34m"

	arg := os.Args[1:]

	files, err := ioutil.ReadDir(arg[0])
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range files {

		switch value.IsDir() {
		case true:
			cl := fmt.Sprintf("%v%v", Blue, value.Name())
			fmt.Printf("%v ", cl)
		case false:
			cl := fmt.Sprintf("%v ", value.Name())
			fmt.Printf("%v", cl)

		}
		fmt.Printf("%v ", value.Name())

	}

	// fmt.Println(Blue, "hello")
}
*/

