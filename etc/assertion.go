// 타입 assertion 에 대해 배웁니다.
package main

import "fmt"

type Animal interface {
	Say(interface{}) string
}

type Person string
func (m Person) Say(s interface{}) string {
	// 타입 assertion
	s.(Person).emit()

	return fmt.Sprintf("Person Say %s %s", m, s)
}

func (m Person) emit() {
	fmt.Println("Oooops...")
}



func main() {
	var animal Animal
	animal = Person("HaHaHa!!")
	fmt.Println(animal.Say(Person("say yayaya!")))
}