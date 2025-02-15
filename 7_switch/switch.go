package main

import (
	"fmt"
	"time"
)

func main() {

	//simple switch
	i := 10
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}

	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	//type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool" )
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Println(t)
		}
	}

	whoAmI(true)
	whoAmI(1)
	whoAmI("golang")
	whoAmI([]int{1, 2, 3})
}
