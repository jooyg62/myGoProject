package prove

import (
	"fmt"
	"strings"
)

func JoinWithCommas(phrases []string) string {
	result := strings.Join(phrases[:len(phrases)-1], ", ")
	result += ", and "
	result += phrases[len(phrases)-1]
	return result
}

func main() {
	fmt.Println(JoinWithCommas([]string{"05", "14", "2018"}))
	fmt.Println(JoinWithCommas([]string{"state", "of", "the", "art"}))
}