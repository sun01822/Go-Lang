package main

import "testing"

func TestDivisibleByThree(t *testing.T){
	type args struct{
		input int 
	}
	tests := []struct{
		name string
		args args
		want string
	}{
		{"TestCase1 for number 9", args{9}, "Hurray!"},
		{"TestCase2 for number 10", args{10}, "10"},
		{"TestCase3 for number 12", args{12}, "Hurray!"},
		{"TestCase4 for number 16", args{16}, "Hurray!"},
	}
	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			if got := DivisibleByThree(tt.args.input); got != tt.want{
				t.Errorf("DivisibleByThree() = %v, want %v", got, tt.want)
			}
		})
	}
}