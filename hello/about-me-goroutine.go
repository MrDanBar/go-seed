package main

import (
	"fmt"
	"time"
)

func main() {

	go hola("HOLA ASYNC")

	hola("HOLA sync")

	time.Sleep(time.Second * 3)

	hola("HOLA done all")
}

func hola(message string) {
	time.Sleep(time.Second)

	fmt.Println(message)
}
