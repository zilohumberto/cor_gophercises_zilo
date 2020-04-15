package main

import "testing"

func TestQuestionWrongResult(t *testing.T) {
	got := question([]string{"2+5", "6"})
	if got == true {
		t.Error("question([]string{'2+5', '6'}) Excepted", false, "got", got)
	}
	got = question([]string{"2*3", "4"})
	if got == true {
		t.Error("question([]string{'2*2', '4'}) Excepted", false, "got", got)
	}
}

func TestQuestionCorrectResult(t *testing.T) {
	got := question([]string{"2+2", "4"})
	if got == false {
		t.Error("question([]string{'2+2', '3'}) Excepted", true, "got", got)
	}
	got = question([]string{"3*5", "15"})
	if got == false {
		t.Error("question([]string{'3*5', '15'}) Excepted", true, "got", got)
	}
}

func TestCalculateSuma(t *testing.T) {
	got := calculate(4, 3, "+")
	if got != 7 {
		t.Error("calculate(4, 3, '+') Excepted", 7, "got", got)
	}
	got = calculate(-3, 3, "+")
	if got != 0 {
		t.Error("calculate(-3, 3, '+') Excepted", 0, "got", got)
	}
}

func TestCalculateResta(t *testing.T) {
	got := calculate(4, 3, "-")
	if got != 1 {
		t.Error("calculate(4, 3, '-') Excepted", 1, "got", got)
	}
	got = calculate(-3, 3, "-")
	if got != -6 {
		t.Error("calculate(-3, 3, '-') Excepted", -6, "got", got)
	}
}

func TestCalculateMulti(t *testing.T) {
	got := calculate(4, 3, "*")
	if got != 12 {
		t.Error("calculate(4, 3, '*') Excepted", 12, "got", got)
	}
	got = calculate(-3, 3, "*")
	if got != -9 {
		t.Error("calculate(-3, 3, '*') Excepted", -9, "got", got)
	}
}
