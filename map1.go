package main

import "fmt"

func main() {
	studentGrade := make(map[string]int) //create map
	studentGrade["aa"] = 100
	studentGrade["bbb"] = 200
	studentGrade["ccc"] = 120
	fmt.Println("marks of aa :-", studentGrade["aa"])
	studentGrade["aa"] = 130 //change value

	fmt.Println("marks of aa :-", studentGrade["aa"])
	delete(studentGrade, "aa") // delet key
	fmt.Println("marks of aa :-", studentGrade["aa"])

	//checking key exist or not

	grade, exist := studentGrade["ddd"]
	fmt.Println("ddd is exist", exist)
	fmt.Println("grade of ddd", grade)

	// using range and for loop
	for index, value := range studentGrade {
		fmt.Printf("key is % s and marks is %d\n", index, value)
	}
	person := map[string]int{
		"ggg": 56,
		"hhh": 67,
		"jjj": 90,
	}
	for index, value := range person {
		fmt.Printf("key is % s and marks is %d\n", index, value)
	}
}
