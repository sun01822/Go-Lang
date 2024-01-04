package main

import (
	"fmt"
)

type Animal interface{
	walk() string
	bark() string
}

type Dog struct{
	name string
	w, b string
	belt string
}

type Cat struct{
	name string
	w, b string
	hat string
}

func (d Dog) walk() string{
	return d.w
}

func (d Dog) bark() string{
	return d.b
}

func (c Cat) walk() string{
	return c.w
}

func (c Cat) bark() string{
	return c.b
}


func main(){
	var a Animal = Dog{b: "Bark Bark.....!!!", w: "Dog is walking.....!!!", belt: "Black"}
	fmt.Println("Animal is: ", a)
	fmt.Println(a.walk())
	fmt.Println(a.bark())
	c:= Cat{b: "Meow Meow.....!!!", w: "Cat is walking.....!!!"}
	fmt.Println(c.walk())
	fmt.Println(c.bark())
}