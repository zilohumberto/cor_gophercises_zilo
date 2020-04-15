// exercise come from: https://github.com/gophercises/quiz
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	// steps
	// read a csv file
	// resolve the problems one problem for line
	// output the total numbers of problem correct and many problems there were in total
	filename := readConsole()
	reader := openFile(filename)
	questions := getLines(reader)
	correct := makeQuestions(questions)
	fmt.Println("count", len(questions), "result", correct)
}

func readConsole() string {
	filename := flag.String("filename", "problem.csv", "CSV File that conatins quiz questions")
	flag.Parse()
	return *filename
}

func openFile(filename string) io.Reader {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return reader
}

type quiz struct {
	question string
	answer   int
}

func getLines(reader io.Reader) []quiz {
	r := csv.NewReader(reader)
	var questions []quiz
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		answer, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}
		questions = append(questions, quiz{record[0], answer})
	}
	return questions
}
func makeQuestions(questions []quiz) int {
	var answerUser int
	var correct int
	for _, v := range questions {
		fmt.Printf("%s ", v.question)
		fmt.Scanf("%d\n", &answerUser)
		if compareResults(v.answer, answerUser) {
			correct++
		}
	}
	return correct
}

func compareResults(resultInput int, resultCalculate int) bool {
	return resultInput == resultCalculate
}
