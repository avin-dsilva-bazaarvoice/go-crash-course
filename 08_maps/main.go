package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign KV
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Declare map and add kv
	emails1 := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	fmt.Println(emails1)
	fmt.Println(emails1["Bob"])
	fmt.Println(len(emails1))

}
