package main

import (
	"fmt"
	"net/http"
)

type escritoWeb struct {
}

func (escritoWeb) Write(p []byte) (int, error) {
	fmt.Println(p)

	return 0, nil
}

func main() {
	respues, _ := http.Get("http://google.com")

	fmt.Println(respues.Body)
}
