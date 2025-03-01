package main

import (
	"fmt"
	"sync"
	// "sync"
)

type post struct {
	views int
	mu sync.Mutex
}

func (p *post) increaseViews(wg *sync.WaitGroup) {
	defer func () {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views++
}
func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for range 100 {
		wg.Add(1)
		go myPost.increaseViews(&wg)
	}
	// myPost.increaseViews()

	wg.Wait() 
	fmt.Println(myPost.views)
}
