package main

import "fmt"

func greeting(name string) string {
	return "Hello, " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}
func main() {
	fmt.Print(greeting("Avin"))
	fmt.Println(getSum(5, 2))
}
