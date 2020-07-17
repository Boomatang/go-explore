package main

import (
	"fmt"
)

type elem struct {
	Name string
}

func main() {
	items := []elem{{Name: "Zero"}, {Name: "One"}, {Name: "Two"}, {Name: "Three"}}

	for index, item := range items {
		worker(index, &item)
	}

}

func worker(index int, item *elem) {
	fmt.Printf("\nindex: %v, element.Name: %s\n", index, item.Name)

}
