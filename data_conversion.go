package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 20
	//num = num + 12.34// error
	fmt.Println("num is:=", num)

	fmt.Printf("type of num is: %T\n", num)
	//num convert to float
	var data float64 = float64(num)
	data = data + 12.54
	fmt.Println("num is", data)
	fmt.Printf("type of num is: %T\n", data)

	//convert int to string
	num = 123
	str := strconv.Itoa(num)
	fmt.Println("str is", str)
	fmt.Printf("type of str is: %T\n", str)
	// convert string to int
	number := "123"
	str1, _ := strconv.Atoi(number)
	fmt.Println("str is", str1)
	fmt.Printf("type of str is: %T\n", str1)
	//convert string to float
	num_string := "123.56"
	num_float, _ := strconv.ParseFloat(num_string, 64)
	fmt.Println("str is", num_float)
	fmt.Printf("type of str is: %T\n", num_float)
}
