package main

import "fmt"

type Book struct {
	name        string
	authorId    string
	timeCreated int64
	tags        []string
}

func (book Book) displayName() string {
	return book.name + ".-:'()"
}

func main() {
	agenda := Book{"My Agenda", "BK-281654891", 168990560510, []string{"COOL", "SAD"}}

	fmt.Println(agenda)
	fmt.Printf("%+v", agenda)
	fmt.Println()

	fmt.Print("name:", agenda.name)
	fmt.Print(" author_id:", agenda.authorId)
	fmt.Print(" book_created_time:", agenda.timeCreated, "\n")

	fmt.Println(agenda.displayName())

	fmt.Println(agenda.tags)

	for _, v := range agenda.tags {
		fmt.Println("NEW TAG ->", v)
	}
}
