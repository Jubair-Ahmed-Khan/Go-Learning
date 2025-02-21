package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func main() {
	// newCustomer := customer{
	// 	name:  "John Doe",
	// 	phone: "123-456-7890",
	// }
	newOrder := order{
		id:     "1",
		amount: 10.99,
		status: "pending",
		// customer: newCustomer,
		customer: customer{
			name:  "John Doe",
			phone: "123-456-7890",
		},
	}
	fmt.Println(newOrder)
}
