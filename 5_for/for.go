package main

import "fmt"

// for -> only construct in go for looping
func main() {

	//while looping using for
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//infinite looping
	// for {
	// 	fmt.Println("Hello")
	// }

	//for with break
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	//for using range
	for i := range 5 {
		fmt.Println(i)
	} 
}
