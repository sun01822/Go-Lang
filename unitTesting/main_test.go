package main

import (
	"testing"
)

func TestDivisibleByThree(t *testing.T){
	got:=DivisibleByThree(9)
	want:="Hurray!"
	if want!=got {
		t.Errorf("Expected '%s', but got '%s'", want, got)
		
	}
}