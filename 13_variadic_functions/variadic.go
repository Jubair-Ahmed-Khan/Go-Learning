package main

import "fmt"

// variadic function
// accepts any number of arguments

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func main() {
	nums := []int{2, 3, 4, 5}
  fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(nums...))
}