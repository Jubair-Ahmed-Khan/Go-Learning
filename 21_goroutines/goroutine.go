package main

import (
	"fmt"
	"sync"
	// "time" 
)

// task performs a task with given id and notifies the wait group when done
func task(id int, w *sync.WaitGroup) {
	defer w.Done() // runs after function finish
	fmt.Println("doing task", id)
}

// main initializes a wait group and starts multiple goroutines to perform tasks concurrently.
// Each goroutine executes the task function, and the main function waits for all goroutines
// to complete before exiting.

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	// time.Sleep(time.Second * 2)
	wg.Wait()
}
