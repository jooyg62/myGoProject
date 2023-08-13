package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.Tick(1 * time.Second)
	count := 0
	for next := range c {
		fmt.Printf("%v", next)
		count += 1
		if count == 10 {
			break
		}
	}
}
