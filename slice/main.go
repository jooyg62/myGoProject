package main

import "fmt"

func main() {
	var mySlice []string
	mySlice = make([]string, 7)
	fmt.Println(mySlice)

	mySlice2 := make([]string, 8)
	mySlice2[3] = "go"
	fmt.Println(mySlice2)

	mySlice2 = append(mySlice2, "a", "b")
	fmt.Println(mySlice2)
}