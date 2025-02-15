package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	fmt.Println("Hello World")
	//closure
	//closure is a function that returns another function
	// inner function can access variables of outer function

	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

}
