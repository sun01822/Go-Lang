package fucntionspointer

import (
	"fmt"
)

func Three(){
	a := 5
	aPtr := &a 
	aPtr2 := &a 

	fmt.Println(a, aPtr, aPtr2)
}