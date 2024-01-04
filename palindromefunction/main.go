package main

import (
	"fmt"
)


func isPalindromePossible(str string) bool{
	len := len(str)
	for i:=0; i<len/2; i++{
		if(str[i] != str[len-i-1]){
			return false
		}
	}
	return true
}

func main(){
	var str string
	fmt.Print("Enter a string: ");
	fmt.Scan(&str)
	if isPalindromePossible(str){
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}