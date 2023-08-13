// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {
	t := tree.Tree{}
	fmt.Println(t)

	if t.Left == nil {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}
}
