package main

import "fmt"

type Person float64
func (m Person) Error() string {
	return fmt.Sprintf("Say Error hahaha: %.2f", m)
}

type Person2 string
func (m Person2) String() string {
	return string(m) + " is person!"
}

func main() {
	var err error = Person(2.0)
	if err != nil {
		fmt.Println(err)
	}

	var p Person2 = Person2("1.0")
	fmt.Println(p)
}