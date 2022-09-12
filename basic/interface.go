package basic

import "fmt"

// Interface
type Animal interface {
	GetName() string
}

func Walk(a Animal) {
	fmt.Println(a.GetName() + " is walking towards you!")
}

// Implement the interface
type Cat struct {
	Name string
}

func (c Cat) GetName() string {
	return c.Name
}

func Interface() {
	motherCat := Cat{"Mother Cat"}
	Walk(motherCat)
}
