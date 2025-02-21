package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

// constructor
func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here ....
	orderDetails := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &orderDetails
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}
func main() {
	// if not assigned, default value is stored
	// int=0, float=0, string="", bool=false

	//constructor
	myOrder := newOrder("1", 10.99, "pending")

	// myOrder := order{
	// 	id:     "1",
	// 	amount: 10.99,
	// 	status: "pending",
	// }

	myOrder.createdAt = time.Now()
	myOrder.changeStatus("processing")

	fmt.Println(*myOrder)
	fmt.Println(myOrder.getAmount())

	//only one instance of struct
	language := struct {
		name   string
		isGood bool
	}{"Go", true}

	fmt.Println(language)

}
