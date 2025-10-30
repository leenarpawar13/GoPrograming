package main

import (
	"fmt"
	"time"
)

func sayHello() {

	fmt.Println("hello everyon")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Say hello function completed")
}
func syahii() {
	fmt.Println("Hii leena")
	time.Sleep(1000 * time.Millisecond)
}
func main() {

	fmt.Println("learning Gorouting")
	go sayHello() // using go keyword func concurrantly run
	go syahii()

	time.Sleep(2000 * time.Millisecond)
}
