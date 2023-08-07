package main

import (
	"fmt"
	"log"
)

func Socialize() error {
	// defer 은 함수가 끝나고 실해이 됩니다.
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	return fmt.Errorf("I don't want to talk.")
	fmt.Println("Nice weather, eh?")
	return nil
}

func main() {
	err := Socialize()
	if err != nil {
		log.Fatal(err)
	}
}