package main

import(
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main(){
	if _, err:= os.Open("text.txt"); err!= nil{
		if errors.Is(err, fs.ErrNotExist){
			fmt.Println("file does not exits")
		}else{
			fmt.Println(err)
		}
	}
}