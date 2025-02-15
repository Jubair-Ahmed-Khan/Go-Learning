package main

import "fmt"

const age = 29

func main() {
	const name = "golang"
	fmt.Println(name)
	fmt.Println(age)

	//constant grouping
	const (
		port = 3000
		host = "localhost"
	)
	fmt.Println(host, ":", port)
}
