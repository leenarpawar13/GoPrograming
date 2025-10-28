package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("web request program...")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") //fake api url
	if err != nil {
		fmt.Println("error getting get response", err)
		return

	}
	defer res.Body.Close()

	//fmt.Println("type of response%t\n", res)
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading get response", err)
		return

	}
	fmt.Println("response", string(data))

}
