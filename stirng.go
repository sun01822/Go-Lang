package main

import "fmt"

func main() {
    str := "Hello, world!"
	num1 := 42
	num2 := 3.14159
	num3 := 1234567890
	boolean := true
	char := 'A'

	// %T: print type of the value
	fmt.Printf("%T\n", str)
    fmt.Printf("%T\n", num1)
    fmt.Printf("%T\n", num2)
    fmt.Printf("%T\n", num3)
    fmt.Printf("%T\n", boolean)
    fmt.Printf("%T\n", char)

	// %v print default format for each type
	fmt.Printf("%v\n", str)
    fmt.Printf("%v\n", num1)
    fmt.Printf("%v\n", num2)
    fmt.Printf("%v\n", num3)
    fmt.Printf("%v\n", boolean)
    fmt.Printf("%v\n", char)


	// string format specifier
	fmt.Printf("%s\n", str)  // %s print string
	fmt.Printf("%q\n", str)  // %q print quoted string
	fmt.Printf("%x\n", []byte(str))  // %x print hex encodingcoding of bytes
	fmt.Printf("%x\n", []byte(str))  // %X print uppercase hex encoding of bytes


	// Integer format specifier
	fmt.Printf("%d\n", num1) // %d print decimal Integer
	fmt.Printf("%b\n", num1) // %b print binary Integer
	fmt.Printf("%o\n", num1) // %o print octal Integer
 	fmt.Printf("%x\n", num1) // %x print hex Integer
	fmt.Printf("%X\n", num1) // %X print hex Integer
	fmt.Printf("%c\n", char) // %c print character


	// Floating-point format specifiers 
	fmt.Printf("%f\n",num2) // %f print floating point number
	fmt.Printf("%e\n",num2) // %e print scientific notation of floating number
	fmt.Printf("%E\n",num2) // %E print scientific notation of floating point number with uppercase E
	fmt.Printf("%g\n",num2) // %g print floating-point number in decimal or scientific notation, depending on the value
	fmt.Printf("%G\n",num2) // %G print floating-point number in decimal or scientific notation, depending on the value

	
	// Width and precision
	fmt.Printf("|%5d|\n", num1)     // %5d print decimal integer with minimum width of 5 characters
	fmt.Printf("|%-5d|\n", num1)    // %-5d print decimal integer with minimum width of 5 characters, left-justified
	fmt.Printf("|%5.2f|\n", num2)   // %5.2f print floating-point number with minimum width of 5 characters and 2 digits after the decimal point
	fmt.Printf("|%-5.2f|\n", num2)  // %-5.2f print floating-point number with minimum width of 5 characters and 2 digits after the decimal point, left-justified

  
	// Boolean format specifier
	fmt.Printf("%t\n", boolean) // %t print boolean value

	// Pointer format specifier
    fmt.Printf("%p\n", &num3)   // %p print pointer address

}