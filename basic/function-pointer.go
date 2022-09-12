package basic

import "fmt"

type Stuff struct {
	Name   string
	Amount int
}

func RestockStuff1(s Stuff) {
	s.Amount = 100
	fmt.Println("Restock1: ", s)
}

func RestockStuff2(s *Stuff) {
	s.Amount = 150
	fmt.Println("Restock1: ", s)
}

func FunctionPointer() {
	// Passing to function not by reference
	stuff1 := Stuff{"Donuts", 0}
	RestockStuff1(stuff1)
	fmt.Println(stuff1)

	// Passing to function by reference
	stuff2 := Stuff{"Macaron", 1}
	RestockStuff2(&stuff2)
	fmt.Println(stuff2)

	stuff3 := &stuff1
	RestockStuff2(stuff3)
	fmt.Println(stuff3)
}
