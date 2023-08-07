package main

import "fmt"

type Person float64
func (m Person) Error() string {
	return fmt.Sprintf("Say Error hahaha: %.2f", m)
}

func main() {
	var err error = Person(2.0, 1.0)
	if err != nil {
		fmt.Println(err)
	}
}