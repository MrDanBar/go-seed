package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "Rojo"
	colors[1] = "Verde"
	colors[2] = "Azul"

	fmt.Println(colors)

	var numbers []string

	numbers = append(numbers, "ichi", "ni", "san")

	fmt.Println(numbers)

	numbers = append(numbers, numbers[1:]...)

	fmt.Println(numbers)

	rgb := make(map[string]string)

	rgb["R"] = "#FF0000"
	rgb["G"] = "#00FF00"
	rgb["B"] = "#0000FF"

	fmt.Println(rgb)
}
