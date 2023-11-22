package main

import(
	"fmt"
)

type Animal interface{
	walk() string
	bark() string
}

type Dog struct {
	w, b string
}

func (d Dog) walk() string {
	return d.w
}

func (d Dog) bark() string {
	return d.b
}


func main() {
	var d Dog 
	d.b = "Hello bark!"
	d.w = "Hello!"
	var a Animal = Dog{"Dog is walking......!!!", "Dog is barking.....!!!"}
	fmt.Println("!..........Dog........!")
	fmt.Println(a.walk())
	fmt.Println(a.bark())
	
	// assign d values to the a Animal
	a = d 
	fmt.Println("!..........Dog........!")
	fmt.Println(a.walk())
	fmt.Println(a.bark())


	b := d 
	fmt.Println(b)
	fmt.Println(b.walk())
	fmt.Println(b.bark())

	b.b = "Changed bark by value"
	b.w = "Changed walk by value"
	fmt.Println("Print by dot: ",b.b)
	fmt.Println("Print by dot: ",b.w)
	fmt.Println("Print by method: ",b.bark())
	fmt.Println("Print by method: ",b.walk())

}