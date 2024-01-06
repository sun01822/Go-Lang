package main

import (
	"fmt"
	"time"
)

// --- sending data into a channel ---
func sendData(ch chan string){
	fmt.Println("Sending data into channel...")

	ch <- "Hello"
	fmt.Println("String has been received from channel")
}

// --- receiving data from a channel ---
func getData(ch chan string){
	time.Sleep(2 * time.Second)
	fmt.Println("String retrived from channel: ", <-ch)
}

func main(){
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	fmt.Scanln()
	fmt.Println("End of main")
}
