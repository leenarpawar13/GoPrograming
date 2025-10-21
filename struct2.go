package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
} // create struct

type contact struct {
	Email string
	phone string
}
type address struct {
	house int
	area  string
	state string
}
type Employee struct {
	person_detail  person
	person_address address
	person_contact contact
}

func main() {
	var employee1 Employee
	employee1.person_detail = person{
		fname: "hii",
		lname: "hello",
		age:   23,
	}
	employee1.person_contact.Email = "abc@gmain.com"
	employee1.person_contact.phone = "123245456"
	fmt.Println("employee detail", employee1)
}
