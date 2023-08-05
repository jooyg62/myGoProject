package main

import (
	"fmt"
	"github.com/headfirstgo/count/datafile"
	"log"
	"sort"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	count := map[string]int{}
	for _, line := range lines {
		count[line]++
	}

	var names []string
	for name, _ := range count {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s: %d\n", name, count[name])
	}

	
}