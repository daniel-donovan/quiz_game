// implements exercise 1 from 'gophersizes'
// https://courses.calhoun.io/lessons/les_goph_01
package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var fileName string
var timerDuration, totalQuestions, rightAnswers int

func ParseCommandLineFlags() {
	flag.StringVar(&fileName, "csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.IntVar(&timerDuration, "limit", 30, "the time limit for each question in seconds")

	flag.Parse()
}

func ReadQuizCsvFile() [][] string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal("Error while reading the file: ", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading problems")
	}

	return records
}

func readAnswerWithTimeLimit(input chan string) {
	for {
		in := bufio.NewReader(os.Stdin)
		result, _ := in.ReadString('\n')

		input <- result
	}
}

func AdministerQuiz(questions [][] string) {
	input := make(chan string, 1)
	go readAnswerWithTimeLimit(input)

	for _, question := range questions {

		fmt.Printf("%s=?\n", question[0])

		select {
		case answer := <- input:
			if question[1] == strings.TrimSpace(answer) {
				rightAnswers++
			}
		case <-time.After((time.Duration(timerDuration) * time.Second)):
			fmt.Print("\nTime's up!\n\n")
			return
		}
	}
}

func main() {
	ParseCommandLineFlags()

	questions := ReadQuizCsvFile()
	totalQuestions := len(questions)

	AdministerQuiz(questions)

	fmt.Printf("You scored %d out of %d", rightAnswers, totalQuestions)
}