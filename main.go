package main

import "fmt"

func main() {
	idGenerator, err := NewIdGenerator(1)
	if err != nil {
		panic(fmt.Sprintf("Cannot init Id gen %v", err))
	}
	for i := 0; i < 10; i++ {
		newId, err := idGenerator.nextId()
		if err != nil {
			panic(fmt.Sprintf("Cannot generate new id %v", err))
		}
		fmt.Println("res", newId)
	}
}
