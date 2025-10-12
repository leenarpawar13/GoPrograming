package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("what is you name")
	//var name string
	reder := bufio.NewReader(os.Stdin) // bufio is package for read input
	name, _ := reder.ReadString('\n')
	fmt.Println("hello mr. ", name)
}
