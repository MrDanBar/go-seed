package main

import (
	"fmt"
	"net/http"
	"time"
)

var saludos = []string{"konichiwa", "101", "hi", "hola", "ciao", "salute"}

func main() {
	http.HandleFunc("/v1/test", MyTest)

	fmt.Println("Listening on port 8080...")

	http.ListenAndServe(":8080", nil)
}

func MyTest(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "My first test (CTF)", "More", "Text")

	channel := make(chan string)

	go Saludar(channel)

	fmt.Fprintf(writer, "====")

	for saludado := range channel {
		fmt.Printf(saludado)
		fmt.Fprintf(writer, saludado)
	}

	fmt.Fprintf(writer, "...")
}

func Saludar(channel chan<- string) {
	for _, value := range saludos {
		time.Sleep(1 * time.Second)
		channel <- value
	}

	close(channel)
}
