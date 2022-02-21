package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "Male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}
func main() {
	// Init person using struct
	person1 := Person{firstName: "Brad", lastName: "Traversey", city: "New York", gender: "Male", age: 25}
	fmt.Println(person1)

	// Alternative
	person2 := Person{"Samantha", "Prabhu", "New York", "Female", 25}
	fmt.Println(person2)
	fmt.Println(person1.firstName)

	person1.age++
	fmt.Println(person1)

	person1.hasBirthday()
	fmt.Println(person1.greet())

	person2.getMarried("Andreson")
	fmt.Println(person2)

}
