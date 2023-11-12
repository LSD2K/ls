package main

import "fmt"

// This function prints the user's information.
func myinfo() string {
	// Get the user's name.
	name := "John Doe"

	// Get the user's age.
	age := 30

	// Get the user's gender.
	gender := "male"

	// Return a string with the user's information.
	return fmt.Sprintf("Name: %s\nAge: %d\nGender: %s", name, age, gender)
}

// The main function prints the user's information.
func main() {
	// Print the user's information.
	fmt.Println(myinfo())
}
