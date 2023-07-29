// 포인터를 이용하여 값을 두배로 만들어주는 함수를 만들어봅니다.
package main

import "fmt"

func double(number *int) {
	*number *= 2
}

func main() {
	number := 2
	double(&number)
	fmt.Printf("%d", number)
}
