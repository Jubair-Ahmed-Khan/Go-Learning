package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	fmt.Println("paid using razorpay", amount)
}

type stripe struct {
}

func (r stripe) pay(amount float32) {
	fmt.Println("paid using stripe", amount)
}

type fakePay struct {
}

func (r fakePay) pay(amount float32) {
	fmt.Println("paid using fakepay", amount)
}

func main() {

	razorpayPaymentGw := razorpay{}
	stripePaymentGw := razorpay{}
	fakePaymentGw := razorpay{}

	payment1 := payment{
		gateway: razorpayPaymentGw,
	}

	payment2 := payment{
		gateway: stripePaymentGw,
	}

	payment3 := payment{
		gateway: fakePaymentGw,
	}

	payment1.makePayment(100)
	payment2.makePayment(200)
	payment3.makePayment(300)
}
