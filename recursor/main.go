// 재귀함수를 호출합니다.
package main

import "fmt"

func count(start int, end int) {
	defer fmt.Println("run before panic!")
	fmt.Println(start)
	if start < end {
		panic("this is panic!!")
		count(start+1, end)
	}
}

func main() {
	count(1, 3)
}