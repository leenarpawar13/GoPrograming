package main

import "fmt"

func main() {
	var name [5]string // declared array
	name[0] = "hii"    //assign values
	name[1] = "hasd"
	name[2] = "sddf"
	fmt.Println("name is =", name)
	number := []int{1, 2, 3, 4, 5, 6, 7, 8} //diff way to declared array and asigning value

	fmt.Println("number is:=", number)
	fmt.Println("number is:=", len(number))
	fmt.Println("value of 2nd index:=", number[2])
}
