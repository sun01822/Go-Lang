package fucntionspointer

import(
	"fmt"
)

func Four(){
	a := 5
	add := func(aParam, b int){
		aParam = aParam + b
		return
	}
	add(a, 2)
	fmt.Println("a: ", a)
}