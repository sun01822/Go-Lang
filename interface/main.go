package main

import (
	"interface/method"
)

func main() {
	mystring := method.MyString("My name is MD. Shariar Hossain Sun")
	rectangular := method.Rectangular{Width:5.2, Height: 7}

	method.EmptyInterface(mystring)
	method.EmptyInterface(rectangular)
}