package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Quiz struct {
	question, answer string
}


func main() {
	nflag := flag.Int("num", 5, "this specifies the number of questions you wish to answer")
	csvFile := flag.String("csv", "problems.csv", "this is used to specify the path of your csv file")
	timeLimit := flag.Int("dur", 5 ,"this specifies the duration of the quiz in seconds")
	flag.Parse()

	quizapp(*nflag, *csvFile, timeLimit)
}




func quizapp(num int, csvFile string, timeLimit *int) {

	var correct int
	fileNotFound := errors.New("oops!, file not found")

	file, err := os.Open(csvFile)

	if err != nil {
		log.Fatal(fileNotFound)
	}

	fileData, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println(err)
	}


	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second) // timer


	if len(fileData) >= num {
		for _, data := range fileData[:num] {

			quiz := Quiz{
				question: data[0],
				answer:   data[1],
			}

			fmt.Printf("%s ", quiz.question)
			answerCh := make(chan string)

			go func(){
				var input string
				fmt.Scanf("%s", &input)
				answerCh <- input
			}()

			select {

			case <-timer.C:
				fmt.Printf("\nyou scored %d out of %d\n", correct, len(fileData[:num]))
				return
			case answer := <- answerCh:
				if answer == quiz.answer {
					correct += 1
				}	
			}

		}

		fmt.Printf("you scored %d out of %d\n", correct, len(fileData[:num]))

	} else {
		log.Fatal("how many questions you tryna answer? :(")

	}

}
