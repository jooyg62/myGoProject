package main

import (
	"fmt"
	"github.com/headfirstgo/testing/prove"
)

func main() {
	fmt.Println(prove.JoinWithCommas([]string{"05", "14", "2018"}))
	fmt.Println(prove.JoinWithCommas([]string{"state", "of", "the", "art"}))
}