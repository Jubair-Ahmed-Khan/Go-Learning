package main

import "fmt"

// func printIntSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }
// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//function generics
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stackInt struct {
	elements []int
}
type stackString struct {
	elements []string
}

//struct generics
// any, comparable, interface{}
type stack[T any] struct {
	elements []T
}

func main() {
	//nums := []int{1, 2, 3}
	// printIntSlice(nums)

	// names := []string{"golang", "typescript"}
	// printStringSlice(names)

	// numbers := []int{2, 3, 5}
	// printSlice(numbers)

	// flowers := []string{"rose", "jesmin", "beli"}
	// printSlice(flowers)

	// myIntStack := stackInt{
	// 	elements: []int{1, 2, 3},
	// }

	// fmt.Println(myIntStack)

	// myStringStack := stackString{
	// 	elements: []string{"aa", "bb", "cc"},
	// }

	// fmt.Println(myStringStack)

	intStack := stack[int] {
		elements: []int{5, 7, 2},
	}

	fmt.Println(intStack)

	stringStack := stack[string] {
		elements: []string{"muhit", "rahul", "shuvo"},
	}
	
	fmt.Println(stringStack)
}
