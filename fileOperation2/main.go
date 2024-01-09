package main

import (
	"fmt"
	"os"
)

func createFile(){
	f, err := os.Create("F:\\GO Lang\\fileOperation2\\test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File Created Successfully")
	defer f.Close()
}

func main(){
	// Remove file
	err := os.Remove("F:\\GO Lang\\fileOperation2\\test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File Removed Successfully")
	
}
