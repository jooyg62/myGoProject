package main

import "fmt"

func main() {
	var notes [5]string = [5]string{
		"a", 
		"b", 
		"c", 
		"d", 
		"e",
	}
	fmt.Printf("%#v\n", notes)
	fmt.Printf("%#v\n", notes[1])
	fmt.Printf("%d\n", len(notes))

	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}
}