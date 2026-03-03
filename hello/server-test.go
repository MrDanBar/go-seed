package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var saludos = []string{"konichiwa", "101", "hi", "hola", "ciao", "salute"}

const WorkerNuber = 5

func main() {
	http.HandleFunc("/v1/tests", MyTests)
	http.HandleFunc("/v1/tests/{id}", MyTestById)

	fmt.Println("Listening on port 8080...")

	http.ListenAndServe(":8080", nil)
}

func MyTests(writer http.ResponseWriter, request *http.Request) {
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

func MyTestById(writer http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(writer, request.Pattern)

	// writer.WriteHeader(404)
	// writer.Write([]byte("H G O L A"))

	myFirstTest := Test{"T-5195a1sd9", "mrdanbar", 178035106510, []string{"Astrid san wa doko desu ka"}}

	writer.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(writer).Encode(myFirstTest)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

type Test struct {
	Id          string   `json:"id"`
	AuthorId    string   `json:"author_id"`
	TimeCreated int64    `json:"time_created"`
	Questions   []string `json:"questions"`
}
