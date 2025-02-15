package main

import (
	"fmt"
	"slices"
)

// slice - dynamic array
// most used constructs in go
// + useful methods
func main() {
	//uninitialized slice is nil

	// var nums []int
	// fmt.Println(nums)
	// fmt.Println(len(nums))

	// if we want not nil slice, we need to make it with make function

	var nums2 = make([]int, 0, 4) // 0 is initial size, 4 is capacity(dynamic)
	fmt.Println(nums2)
	fmt.Println("length:", len(nums2))
	fmt.Println("capacity:", cap(nums2))

	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 3)
	nums2 = append(nums2, 4)
	nums2 = append(nums2, 5)
	fmt.Println(nums2)
	fmt.Println("length:", len(nums2))
	fmt.Println("capacity:", cap(nums2))

	//alternative way to create slice
	nums3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums3)

	//copy slice
	var arr1 = make([]int, 0, 5)
	arr1 = append(arr1, 2)

	var arr2 = make([]int, len(arr1))

	copy(arr2, arr1)

	fmt.Println(arr1, "=", arr2)

	//slice operator
	var nums4 = []int{1, 2, 3}
	fmt.Println(nums4[0:1])
	fmt.Println(nums4[:1])
	fmt.Println(nums4[1:])

	//slices package methods
	var a1 = []int{1, 2}
	var a2 = []int{1, 2}

	fmt.Println("Is equal =", slices.Equal(a1, a2))
	fmt.Println("clone =", slices.Clone(a2))

	//2d slice
	var a3 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2d slice =", a3)

}
