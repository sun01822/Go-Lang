package main

import (
	"fmt"
)

func recursion(n int)int{
	if(n==1){
		return 1
	}else{
		return n * recursion(n-1) 
	}
}

func main() {
	var n int 
	fmt.Print("Enter the number: ");
	fmt.Scan(&n)
	ans := recursion(n)
	fmt.Println("The factorial of ", n, " is ", ans)
}