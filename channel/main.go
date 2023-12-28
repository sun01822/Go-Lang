package main

import(
	"fmt"
)

func main(){
	numberCh := make(chan int, 10)
	for i:=1 ; i<=10; i++ {
		numberCh <- i
	}
	for i:=1 ; i<=10 ; i++ {
		fmt.Println(<-numberCh)
	}

}