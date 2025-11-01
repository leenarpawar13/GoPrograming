package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processnum(numchan chan int) {

	for num := range numchan {
		fmt.Println("recived data ", num)
		time.Sleep(time.Second)

	}

}

func main() {

	numchan := make(chan int)

	go processnum(numchan)
	for {

		numchan <- rand.Intn(100)
	}

	time.Sleep(time.Second * 2)

	/*fmt.Println("learnung channels")
	msgchannels := make(chan string) // create channels
	msgchannels <- "helllo"          // send data
	msg := <-msgchannels             //recived data

	fmt.Println("chahhels", msg)*/

}
