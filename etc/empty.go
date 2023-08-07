package main

import "fmt"

// go 1.18 버전 부터는 interface{} 를 any 로 사용할 수 있습니다.
func say(a any) {
	fmt.Println(a)
}

func main() {
	say("abc")
	say(1.24)
	say(5)
	say(false)
}