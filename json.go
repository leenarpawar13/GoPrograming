package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Isadult bool   `json:"is_adult"`
}

func main() {

	fmt.Println("json data..")
	person := Person{Name: "leena", Age: 25, Isadult: true}
	fmt.Println("person data", person)
	//convert data to json data encoding(marshalling)
	jsondata, err := json.Marshal(person)
	if err != nil {

		fmt.Println("error", err)
		return
	}
	fmt.Println("json data is", string(jsondata))

	var persondata Person
	err = json.Unmarshal(jsondata, &persondata)
	if err != nil {
		fmt.Println("error", err)
		return

	}
	fmt.Println("json data is", persondata)

}
