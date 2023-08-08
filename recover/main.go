package main

import "fmt"

func calmDown() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	} else {
		panic(p)
	}
	
}

func freakOut() {
	defer calmDown()
	panic("string panic")
	err := fmt.Errorf("there's an error")
	panic(err)
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}