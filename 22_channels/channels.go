package main

import (
	"fmt"
	// "time"
)

// "math/rand"

// sending data to function from goroutine through channel
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second)
// 	}
// 	// fmt.Println("processing number", <-numChan)
// }

// // sending data from function to goroutine through channel
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult

// }

// goroutine synchronizer
// func task(done chan bool){
// 	defer func() {done <- true}()
// 	fmt.Println("processing task")
// }

// func emailSender(emailChannel chan string, done chan bool) {
// 	defer func() {done <- true}()
// 	for email := range emailChannel {
// 		fmt.Println("email sent to", email)
// 		time.Sleep(time.Second)
// 	}
// }

// receive only channel = emailChan <-chan string
// send only channel = emailChan chan<- string

func main() {
	// channels -> communication between goroutines
	// channels are used to send data from one goroutine to another

	// messageChan := make(chan string)

	// messageChan <- "hello world"

	// msg := <-messageChan

	// fmt.Println(msg)

	// numChan := make(chan int)
	// go processNum(numChan)

	// // numChan <- 10
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// result := make(chan int)
	// go sum(result, 10, 20)

	// res := <-result
	// fmt.Println("Sum = ", res)

	// time.Sleep(time.Second * 2)

	// done := make(chan bool)
	// go task(done)

	// <-done

	// emailChannel := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChannel, done)

	// for i := range 5 {
	// 	emailChannel <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println(("done sending emails"))

	// close(emailChannel)
	// <-done

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "hello"
	}()

	for range 2 {
		select {
		case num := <-chan1:
			fmt.Println("received number from chan1", num)
		case str := <-chan2:
			fmt.Println("received string from chan2", str)
		}
	}




}
