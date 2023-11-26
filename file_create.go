package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// create a file with file name example.txt
	file, err:= os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString("Hello World!")
	// close the file 
	// Write something to example.txt
	content, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}