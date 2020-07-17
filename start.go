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

	for index, item := range items {
		wg.Add(1)
		go worker(index, &item, &wg)
	}
	wg.Wait()
}

func worker(index int, item *elem, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("\nindex: %v, element.Name: %s\n", index, item.Name)

}
