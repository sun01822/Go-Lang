package main

import ("fmt")

func main() {
	fmt.Println("Hello World!")

	value := display()
	fmt.Println(value)
}

func display() string{
	return "Hello World!"
}
