package main

import "fmt"

// numbered sequence of specific length of same type
func main() {
	// if not assigned, default value is stored
	var students [5]string

	students[0] = "John"
	students[2] = "Jane"

	fmt.Println(students)

	// var nums [5]int = [5]int{1, 2, 3, 4, 5}
	nums := [5]int{1, 2, 3, 4, 5}

	fmt.Println(len(nums))
	fmt.Println(nums[1])
	fmt.Println(nums)

	//2d Array
	matrix := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(matrix)

	// if size predictable then used
	// Memory Optimization
	// Constant  time access
}
