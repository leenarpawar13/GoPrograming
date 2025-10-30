package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userid"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func PerformeGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") //fake api url
	if err != nil {
		fmt.Println("error getting get response", err)
		return

	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response:", res.Status)
		return
	}

	//fmt.Println("type of response%t\n", res)
	/*data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading get response", err)
		return

	}
	fmt.Println("response", string(data))
	*/
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {

		fmt.Println("error reading get response", err)
		return
	}
	fmt.Println("ToDo", todo)
}
func PerformePostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "heloo",
		Completed: true,
	}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	//convert json data to string
	jsonString := string(jsonData)
	jsonreader := strings.NewReader(jsonString)
	myurl := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myurl, "Application/json", jsonreader)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("post mrthod", string(data))
}
func performePutRequest() {
	todo := Todo{
		UserID:    235,
		Title:     "heloo golang",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("error getting get response", err)
		return

	}
	jsonString := string(jsonData)
	jsonreader := strings.NewReader(jsonString)

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, myurl, jsonreader)
	if err != nil {
		fmt.Println("error getting get response", err)
		return

	}
	req.Header.Set("content-stype", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("err", err)
		return

	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("put mrthod", string(data))
	fmt.Println("status:=", res.Status)
}
func performeDeleteRequest() {

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("error getting get response", err)
		return

	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("err", err)
		return

	}
	defer res.Body.Close()

	fmt.Println("status:=", res.Status)

}

func main() {
	fmt.Println("web request program...")
	//PerformeGetRequest()
	//PerformePostRequest()
	//performePutRequest()
	performeDeleteRequest()
}
