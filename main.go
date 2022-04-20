package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world!!")

	resp, err := http.Get("http://google.com/")

	if err != nil {
		fmt.Println("Noting but Error")
	}
	fmt.Println(resp)
}
