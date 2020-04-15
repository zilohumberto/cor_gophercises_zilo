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
	"time"
)

type quiz struct {
	question string
	answer   int
}

func main() {
	// steps
	// read a csv file
	// resolve the problems one problem for line
	// output the total numbers of problem correct and many problems there were in total
	filename, timeout := readConsole()
	reader := openFile(filename)
	questions := getLines(reader)
	c := make(chan int)
	fmt.Printf("please enter any key to start")
	fmt.Scanf("%v")
	go makeQuestions(questions, c)
	go sleepQuiz(timeout, c)
	correct := readAnswer(c)
	fmt.Println("count", len(questions), "result", correct)
}

func readAnswer(c <-chan int) int {
	var correct int
	for {
		select {
		case _, ok := <-c:
			if !ok {
				return correct
			}
			correct++
		}
	}
}

func readConsole() (string, int) {
	filename := flag.String("filename", "problem.csv", "CSV File that conatins quiz questions")
	timeout := flag.Int("timeout", 30, "Timeout to finish the quiz")
	flag.Parse()
	return *filename, *timeout
}

func openFile(filename string) io.Reader {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return reader
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

func sleepQuiz(timeout int, c chan<- int) {
	boom := time.After(time.Duration(timeout) * time.Second)
	for {
		select {
		case <-boom:
			close(c)
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
func makeQuestions(questions []quiz, c chan<- int) {
	var answerUser int
	for _, v := range questions {
		fmt.Printf("%s ", v.question)
		fmt.Scanf("%d\n", &answerUser)
		if compareResults(v.answer, answerUser) {
			c <- 1
		}
	}
	close(c)
}

func compareResults(resultInput int, resultCalculate int) bool {
	return resultInput == resultCalculate
}
