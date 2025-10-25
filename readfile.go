package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.Open("emp.txt")
	if err != nil {
		fmt.Println("error while writing file")
		return
	}
	defer file.Close()

	//create buffer to read file
	buffer := make([]byte, 1202)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break

		}
		if err != nil {
			fmt.Println("error while read file")

			return
		}
		fmt.Println(string(buffer[:n]))
	}
	//read file into byte a slicw
	content, err := ioutil.ReadFile("emp.txt")
	if err != nil {
		fmt.Println("error while read file", err)

		return
	}
	fmt.Println(string(content))
}
