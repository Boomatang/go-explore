package main

import (
	"fmt"
	"sync"
)

type elem struct {
	Name string
}

func main() {
	var wg sync.WaitGroup
	items := []elem{{Name: "Zero"}, {Name: "One"}, {Name: "Two"}, {Name: "Three"}}

	for index := range items {
		wg.Add(1)
		go worker(index, &items[index], &wg)
	}

	wg.Wait()
	fmt.Println("\nPrinting updated items in the list")
	for index, item := range items {
		fmt.Printf("\nindex: %v, element.Name: %s\n", index, item.Name)
	}
}

func worker(index int, item *elem, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("\nindex: %v, element.Name: %s\n", index, item.Name)
	item.Name = item.Name + " -> Updated"

}
