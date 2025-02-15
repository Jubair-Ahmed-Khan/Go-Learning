package main

import "fmt"

func main() {

	age := 10

	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is a teenager")
	} else {
		fmt.Println("person is a child")
	}

	//alternative - declare variable in if
	if age := 15; age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is a teenager")
	} else {
		fmt.Println("person is a child")
	}

	var role = "admin"
	var hasPermissions = true

	if role == "admin" || hasPermissions {
		fmt.Println("person is an admin")
	}

	//go has no ternary operator

}
