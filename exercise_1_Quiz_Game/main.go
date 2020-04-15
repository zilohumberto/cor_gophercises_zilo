package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func readLines(c chan []string) {
	in := `question, result
3+5,10
5+4,9
3+2,6
`
	r := csv.NewReader(strings.NewReader(in))
	_, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
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
	question := s[0]
	possibleResult, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}
	var operandoLadoA, operandoLadoB, operador string
	operador = ""
	for _, _q := range question {
		q := string(_q)
		if operador != "" {
			operandoLadoB += q
			continue
		}
		if q == "+" || q == "-" {
			operador = q
			continue
		}
		operandoLadoA += q
	}
	operandoLadoAInt, err := strconv.Atoi(operandoLadoA)
	operandoLadoBInt, err := strconv.Atoi(operandoLadoB)
	calculateResult := calculate(operandoLadoAInt, operandoLadoBInt, operador)
	return calculateResult == possibleResult
}

func calculate(a int, b int, o string) int {
	switch o {
	case "+":
		return a + b
	case "-":
		return a - b
	}
	return 0
}

func main() {
	// steps
	// read a csv file
	// resolve the problems one problem for line
	// output the total numbers of problem correct and many problems there were in total
	c := make(chan []string)
	go readLines(c)
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
