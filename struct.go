package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
} // create struct

func main() {
	var abc person
	fmt.Println(" abc data is", abc)
	abc.fname = "hii"
	abc.lname = "hello"
	abc.age = 23
	fmt.Println("assing value", abc)

	person1 := person{ //assing value 2nd way
		fname: "aaa",
		lname: "ggg",
		age:   25,
	}
	fmt.Println("pesron1 detail", person1)

	var person2 = new(person) // using new keyword
	person2.fname = "leena"
	person2.lname = "pawar"
	person2.age = 25
	fmt.Println("person2 details", person2)
	fmt.Println("person2 details", person2.fname)
}
