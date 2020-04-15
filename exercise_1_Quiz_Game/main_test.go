package main

import "testing"

func TestCompareResults(t *testing.T) {
	got := compareResults(3, 6)
	if got != false {
		t.Error("Excepted false, got ", got)
	}
	got = compareResults(3, 3)
	if got != true {
		t.Error("Excepted true, got ", got)
	}
}
