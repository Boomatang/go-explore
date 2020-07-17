package main

import (
	"fmt"
	"sync"
)

type elem struct {
	Name string
}

func main() {
	items := []elem{{Name: "Zero"}, {Name: "One"}, {Name: "Two"}, {Name: "Three"}}

	fmt.Println("\nCall none async functions")
	normal_list(items)

	fmt.Println("\nCall async functions with pointers")
	broken_goroutines(items)

	fmt.Println("\nCall async functions no pointers")
	fixed_goroutines(items)

}

func normal_list(elements []elem) {
	for index, item := range elements {
		worker_pointer(index, &item)
	}
}

func worker_pointer(index int,  item *elem) {
	fmt.Printf("\nindex: %v, element.Name: %s\n", index, item.Name)

}

func broken_goroutines(elements []elem)  {
	var wg sync.WaitGroup
	for index, item := range elements {
		wg.Add(1)
		go async_worker_pointer(index, &item, &wg)
	}
	wg.Wait()
}

func async_worker_pointer(index int,  item *elem, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("\nindex: %v, element.Name: %s\n", index, item.Name)

}

func fixed_goroutines(elements []elem)  {
	var wg sync.WaitGroup
	for index, item := range elements {
		wg.Add(1)
		go async_worker(index, item, &wg)
	}
	wg.Wait()
}

func async_worker(index int,  item elem, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("\nindex: %v, element.Name: %s\n", index, item.Name)

}
