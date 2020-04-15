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
	c := make(chan []string)
	go readLines(reader, c)
	count, correct := 0, 0
	var v []string
	ok := true
	for ok {
		select {
		case v, ok = <-c:
			if !ok {
				continue
			}
			count++
			if question(v) {
				correct++
			}
		}
	}
	fmt.Println("count", count, "result", correct)
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

func readLines(reader io.Reader, c chan []string) {
	r := csv.NewReader(reader)
	for {
		record, err := r.Read()
		if err == io.EOF {
			close(c)
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		c <- record
	}
}

func question(s []string) bool {
	operatorLeftSide, operatorSideRight, operand := parseOperation(s[0])
	possibleResult, err := strconv.Atoi(s[1])
	if err != nil {
		return false
	}
	operatorLeftSideAInt, err := strconv.Atoi(operatorLeftSide)
	if err != nil {
		return false
	}
	operatorRightSideBInt, err := strconv.Atoi(operatorSideRight)
	if err != nil {
		return false
	}
	resultCalculate := calculate(operatorLeftSideAInt, operatorRightSideBInt, operand)
	return compareResults(possibleResult, resultCalculate)
}

func compareResults(resultInput int, resultCalculate int) bool {
	return resultInput == resultCalculate
}

func parseOperation(questions string) (string, string, string) {
	var operatorSideLeft, operatorRightSide, operand string
	for _, _question := range questions {
		q := string(_question)
		if operand != "" {
			operatorRightSide += q
			continue
		}
		if q == "+" || q == "-" || q == "*" {
			operand = q
			continue
		}
		operatorSideLeft += q
	}
	return operatorSideLeft, operatorRightSide, operand
}

func calculate(leftSide int, rightSide int, operand string) int {
	switch operand {
	case "+":
		return leftSide + rightSide
	case "-":
		return leftSide - rightSide
	case "*":
		return leftSide * rightSide
	}
	return 0
}
