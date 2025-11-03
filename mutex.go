package main

import (
	"fmt"
	"sync"
)

type post struct {
	view int
	mu   sync.Mutex //call mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.view += 1
}

func main() {
	var wg sync.WaitGroup
	mypost := post{
		view: 0,
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go mypost.inc(&wg)
	}
	wg.Wait()
	fmt.Println(mypost.view)

}
