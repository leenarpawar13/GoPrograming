package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("current time:=", currentTime)
	formatted := currentTime.Format("2006-01-02,15:04:05, Monday") // time formate in go 2006-01-02 3:04:05 monday
	fmt.Println("time is", formatted)
	formatted = currentTime.Format("2006-01-02,3:04:05 pm, Monday") //fix formate go launched date

	fmt.Println("time is", formatted)

	//convert str to time
	layout_str := "02/01/2006"
	datastr := "16/11/2023"
	formatted_time, _ := time.Parse(layout_str, datastr)
	fmt.Println(formatted_time)

	//add 1 more daynew_data
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println(new_date)
	formatted_new_date := new_date.Format("02/01/2006,Monday")
	fmt.Println("fomatted date", formatted_new_date)
}
