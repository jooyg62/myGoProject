package main

import (
	"fmt"
	"github.com/headfirstgo/struct/magazine"
)

func main() {
	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)

	subscriber := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
	fmt.Println("Name:", subscriber.Name)
	fmt.Println("Rate:", subscriber.Rate)
	fmt.Println("Active:", subscriber.Active)

	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)

	var address magazine.Address
	address.Street = "123 Oak St"
	address.City = "Omaha"
	address.State = "NE"
	address.PostalCode = "68111"
	fmt.Println(address)

	address2 := magazine.Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"}
	subscriber2 := magazine.Subscriber{Name: "Aman Singh"}
	subscriber2.Address = address2
	fmt.Println(subscriber2.Address)
	fmt.Println(subscriber2.City) // 임베딩된 구조체에 접근 가능
}
