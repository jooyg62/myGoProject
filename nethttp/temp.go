package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	reponseSize("https://example.com/")
	reponseSize("https://golang.org/")
	reponseSize("https://golang.org/doc")
}

func reponseSize(url string) {
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
}