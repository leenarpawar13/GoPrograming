package main

import (
	"fmt"
	"strings"
)

func main() {
	//seprate string
	str := "apple,banana,orange"
	data := strings.Split(str, ",")
	fmt.Println("data is:=", data)

	//count word
	str1 := "one ,two,two,three,four,five,two"
	data1 := strings.Count(str1, "two")
	fmt.Println("count is", data1)

	// trim white space
	str2 := "             hello everyone"
	data2 := strings.TrimSpace(str2)

	fmt.Println(data2)

	//join two string
	str1 = "leena"
	str2 = "pawar"
	data3 := strings.Join([]string{str1, str2}, " ")
	fmt.Println(data3)

}
