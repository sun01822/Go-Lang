package main

import (
	"fmt"
)

func isPalindromePossible(str string) bool {
	for i:=0; i<len(str)/2; i++ {
		if str[i] != str[len(str)-i-1]{
			return false
		}
	}
	return true
}

func main(){
	var str string
	fmt.Printf("Enter a string: ")
	fmt.Scan(&str)

	if isPalindromePossible(str) {
		fmt.Printf("%s is a palindrome\n", str)
	}else{
		fmt.Printf("%s is not a palindrome\n", str)
	}

}