package main

import "fmt"

func main() {
	var mySlice []chan string
	mySlice = make([]chan string, 7)
	for i := range mySlice {
		mySlice[i] = make(chan string)
	}
	fmt.Println(mySlice)

	mySlice2 := make([]string, 8)
	for i := range mySlice2 {
		mySlice2[i] = "7"
	}
	mySlice2[3] = "go"
	fmt.Println(mySlice2)

	mySlice2 = append(mySlice2, "a", "b")
	fmt.Println(mySlice2)
}
