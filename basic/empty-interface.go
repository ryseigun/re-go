package basic

import "fmt"

// The interface type that specifies zero methods
// Empty interfaces are used by code that handles values of unknown type
func Ups(i int) interface{} {
	if i == 1 {
		return 2
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func EmptyInterface() {
	var result interface{}
	result = Ups(1)

	fmt.Println(result)
}
