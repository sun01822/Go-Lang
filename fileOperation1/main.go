package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func createFile(){
	// Create a file using Create() function and filePath
	filePath := "f:\\GO Lang\\fileOperation1\\text.txt"
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println("File opened successfully")
	}
	defer f.Close()
}


func openFile(){
	// Open a file using Open() function and filePath
	filePath := "f:\\GO Lang\\fileOperation1\\example.txt"
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println("File opened successfully")
		data := make([]byte, 100)
		count, err := f.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Bytes read: ", count)
		fmt.Println("String read: ", string(data))
	}
	defer f.Close()
}

func readFile(){
	ioutilData, err := ioutil.ReadFile("f:\\GO Lang\\fileOperation1\\example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println("File opened successfully")
		fmt.Println("String read: ", string(ioutilData))
	}
}

func writeFile(){
	f, err := os.Create("f:\\GO Lang\\fileOperation1\\example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println("File opened successfully")
	}
	defer f.Close()
	// Write data to file
	data := []byte("Hello World!")
	count, err := f.Write(data)
	if len(data) != count {
		fmt.Printf("Line num not matched %v, %v\n", len(data), count)
		return 
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(count, "bytes written successfully")

}


func main(){
	//createFile()
	//openFile()
	//readFile()
	writeFile()
}