package main

import "fmt"

type myType string

func (m myType) sayHi() {
	fmt.Println("hi", m)
}

func main() {
	value := myType("a MyType Value")
	value.sayHi()
	anotherValue := myType("another value")
	anotherValue.sayHi()
}