package basic

import "fmt"

// Defining return value by variable name
func GetFullName() (firstName, middleName, lastName string) {
	firstName = "A"
	middleName = "B"
	lastName = "C"

	return // Automatically return the defined variables
}

func NamedReturn() {
	a, b, _ := GetFullName()
	fmt.Println(a, b)
}
