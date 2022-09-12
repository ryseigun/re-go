package basic

import (
	"fmt"
)

func Variable() {
	// How to make variable
	var name string
	name = "Abc"

	fmt.Println(name)

	// Or better
	name2 := "Xyz"
	fmt.Println(name2)

	// Multiple variable declaration
	var (
		firstName = "Shinji"
		lastName  = "Kido"
	)
	fmt.Println(firstName, lastName)

	// Constant declaration
	const riderName string = "Ryuki"

	fmt.Println(riderName)

	// Multiple constant declaration
	const (
		hostName  = "Ikki"
		demonName = "Vice"
		year      = 2021
	)
	fmt.Println(hostName, demonName)
}
