package main

import (
	"fmt"
	"net/url"
)

func main() {

	fmt.Println("handle url")
	myUrl := "https://github.com4353/leenarpawar13/GoPrograming"

	fmt.Printf("urls type is %T\n", myUrl)
	parseurl, err := url.Parse(myUrl)

	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Printf("urls type is %T\n", parseurl)

	fmt.Println("scheme:=", parseurl.Scheme)
	fmt.Println("host:=", parseurl.Host)
	fmt.Println("path:=", parseurl.Path)
	fmt.Println("rawquery:=", parseurl.RawQuery)

	//modified url
	parseurl.Path = "/linarpawar29"
	newurl := parseurl.String() //cpnnect url string into url onject

	fmt.Println("new url", newurl)
}
