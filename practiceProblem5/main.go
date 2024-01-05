package main

import (
	"fmt"
)

type Employee struct{
	name string
	age int
	salary float64
}

func (e *Employee) updateSalary(){
	if(e.salary > 10000){
		e.salary += e.salary * .20 
	}else{
		e.salary += e.salary * .10
	}
}


func main(){
	emp1 := Employee{"Sun", 30, 12000}
	fmt.Println("Before update: ", emp1)
	emp1.updateSalary()
	fmt.Println("After update: ", emp1)
}