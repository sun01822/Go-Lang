package main

import "testing"

func TestGreaterThan10(t *testing.T) {
	got := greaterThan10(11)
	want := "ZZZ"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDivideby3(t *testing.T) {
	got := divideby3(9)
	want := "Yes"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
