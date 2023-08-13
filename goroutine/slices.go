// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	s := []string{"a", "b", "c", "d"}
	fmt.Println(slices.Contains(s, "a"))
}
