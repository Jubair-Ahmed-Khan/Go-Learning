package main

import "fmt"

// range is used to iterate over a collection
func main() {
	nums := []int{6, 7, 8}

	//normal for loop
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	sum := 0
	//range with slice 
	for i, num := range nums {
		fmt.Println(i, num)
		sum = sum + num
	}

	fmt.Println("sum:", sum)

	//range with map
	m := map[string]string{"fname":"john", "lname":"doe"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	//range with string
	// starting byte of rune i
	// unicode code point rune c
	for i, c := range "hello" {
		// fmt.Println(i, c)
		fmt.Println(i,string(c))
	}
}
