package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type Quiz struct {
	question string
	answer string
}

func main(){
	var inp string
	var correct int 
	var wrong int
	
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opening csv csvFile...")
	defer csvFile.Close()

	csvLine, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	quest := flag.Int("num", 5, "specifies the number of questions you wish to answer")
	flag.Parse()

	for _, val := range csvLine[0:*quest] {
		quiz := Quiz{
			question: val[0],
			answer: val[1],
		}
		fmt.Printf("%v = ",quiz.question)
		fmt.Scanln(&inp)

		if inp == quiz.answer {
			correct += 1
		}
		if inp != quiz.answer {
			wrong += 1
		}

	}
	fmt.Printf("correct: %d\nwrong: %d\n",correct,wrong)
}