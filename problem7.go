package main

import (
	"fmt"
)

func main() {
	for i:=1; i<=6; i++ {
		for j:=1; j<=i*2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}	
}