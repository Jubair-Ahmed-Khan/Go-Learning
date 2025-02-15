package main

import "fmt"

func changeNumber(num int) {
	num = 5
	fmt.Println(num)
}

func changeNumber2(num *int) {
	*num = 7
	fmt.Println(*num)
}

func main() {

	n1 := 4
	changeNumber(n1)
	fmt.Println(n1)

	n2 := 4
	changeNumber2(&n2)
	fmt.Println(n2)
}
