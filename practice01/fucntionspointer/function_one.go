package fucntionspointer

import	(
	"fmt"
)

func One() {
	var a int
	var aPtr *int

	a = 5
	aPtr = &a

	fmt.Println(a, &a, aPtr, *aPtr)
}