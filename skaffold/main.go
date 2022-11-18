package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	request1()
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to the Go Web API!")
	fmt.Println("Endpoint Hit: homePage")
}

func request1() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
