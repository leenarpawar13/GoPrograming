package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// foe creating file
	file, err := os.Create("emp.txt")
	if err != nil {
		fmt.Println("error ehile creting file")
		return
	}
	defer file.Close()

	fmt.Println("sucessfully create file")

	// for add contain to file

	contain := "hello..this is go programing"
	_, err1 := io.WriteString(file, contain+"\n")
	if err1 != nil {
		fmt.Println("error while writing file")
		return
	}
	//defer file.Close()

	fmt.Println("sucessfully write file")

}
