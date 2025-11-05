package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	oddch := make(chan bool)
	evench := make(chan bool)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {

			<-oddch
			fmt.Println("odd", i)
			evench <- true
		}

	}()

	go func() {
		defer wg.Done()

		for i := 2; i <= 10; i += 2 {
			<-evench
			fmt.Println("Even", i)
			if i < 10 {
				oddch <- true
			} else {
				close(oddch)
				close(evench)
			}

		}
	}()
	oddch <- true
	wg.Wait()
}
