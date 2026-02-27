package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var text string = "HOLA MUNDO"
	const text2 string = "AIND"
	text3 := "."

	fmt.Println(text, text2, text3)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a text: ")

	value, _ := reader.ReadString('\n')

	fmt.Println(value)

	number, error := strconv.ParseFloat(value, 64)

	if error != nil {
		myError := fmt.Errorf("Houston tenemos problemas.. %v", value)
		fmt.Println("e:", myError)
	}

	fmt.Println("Number is:", number)
}
