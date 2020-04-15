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
	questions := s[0]
	possibleResult, err := strconv.Atoi(s[1])
	if err != nil {
		return false
	}
	var operandoLadoA, operandoLadoB, operador string
	operador = ""
	for _, _question := range questions {
		q := string(_question)
		if operador != "" {
			operandoLadoB += q
			continue
		}
		if q == "+" || q == "-" || q == "*" {
			operador = q
			continue
		}
		operandoLadoA += q
	}
	operandoLadoAInt, err := strconv.Atoi(operandoLadoA)
	if err != nil {
		return false
	}
	operandoLadoBInt, err := strconv.Atoi(operandoLadoB)
	if err != nil {
		return false
	}
	calculateResult := calculate(operandoLadoAInt, operandoLadoBInt, operador)
	return calculateResult == possibleResult
}

func calculate(a int, b int, o string) int {
	switch o {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	}
	return 0
}
