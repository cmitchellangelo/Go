package main

import "fmt"

func main() {

	// Define map

	// emails := make(map[string]string)

	// // Asssign Key Values

	// emails["Bob"] = "bob@gmail.com"
	// emails["Chase"] = "chase@gmail.com"
	// emails["Mike"] = "Mike@gmail.com"

	// Declare map and add key values

	emails := map[string]string{"Bob": "bob@gmail.com", "Chase": "Chase@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete from Map

	delete(emails, "Bob")
	fmt.Println(emails)
}
