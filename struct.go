package main

import "fmt"

type Person struct {
	name string
	age int
	job string
	salary float64
}

func main(){
	var person1 Person
	var person2 Person

	// Person1 specification
	person1.name = "Sun"
    person1.age = 25
    person1.job = "Software Engineer"
    person1.salary = 100000.030

    // Person2 specification
    person2.name = "Moon"
    person2.age = 22
    person2.job = "Vetenary Doctor"
    person2.salary = 23450000.029492

    fmt.Println(person1)
    fmt.Println(person2)

	// Access and print Person1 
	fmt.Println("Name: ", person1.name)
	fmt.Println("Age: ", person1.age)
	fmt.Println("Salary: ", person1.salary)
	fmt.Println("Job: ", person1.job + "\n")


	// Access and print Person2
	fmt.Println("Name: ", person2.name)
    fmt.Println("Age: ", person2.age)
    fmt.Println("Salary: ", person2.salary)
    fmt.Println("Job: ", person2.job)


	// Print Person1 info by calling a function
	printPerson(person1)


	// Print Person2 info by calling a function
	printPerson(person2)


}


func printPerson(person Person) {
	fmt.Println("\nBy calling a function: ")
	fmt.Println("Name: ", person.name)
	fmt.Println("Age: ", person.age)
	fmt.Println("Salary: ", person.salary)
	fmt.Println("Job: ", person.job + "\n")
}