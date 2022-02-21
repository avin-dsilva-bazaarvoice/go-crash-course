package main

import "fmt"

// NOTE : Declaration short hand not allowed outside of a function
// ./main.go:5:1: syntax error: non-declaration statement outside function body
// name := "Milind"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float 32 float64
	// complex64 complex128

	// Using var
	var name string = "Milind" // string is infered, you don't need to put it
	var age int32 = 50         // implicit typing
	const isCool = true

	// Shorthand
	surname := "Soman"
	email := "milind@theinternet.com"

	height, weight := 173, 73

	fmt.Println(name, surname, email, age, isCool, height, weight)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)

}
