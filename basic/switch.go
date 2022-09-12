package basic

import "fmt"

func Switch() {
	// How to make switch 1
	name := "Kido"

	switch name {
	case "Godai":
		fmt.Println("Kuuga")
	case "Kido":
		fmt.Println("Ryuki")
		fmt.Println("Knight")
	case "Kiryu":
		fmt.Println("Build")
	default:
		fmt.Println("Decade")
	}

	// How to make switch 2
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}
}
