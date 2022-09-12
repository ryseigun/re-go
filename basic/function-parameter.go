package basic

import (
	"fmt"
	"strings"
)

type filter func(string) string

func SayHello(name string, f filter) string {
	return "Hello, " + f(name)
}

func Filter1(name string) string {
	if name == "Anjing" {
		return "..."
	}

	return name
}

func Filter2(name string) string {
	return strings.ToUpper(name)
}

func FunctionParameter() {
	fmt.Println(SayHello("Fourze", Filter2))
	fmt.Println(SayHello("Anjing", Filter1))

	// Anonymous function
	fmt.Println(SayHello("GEATS", func(name string) string {
		return strings.ToLower(name)
	}))
}
