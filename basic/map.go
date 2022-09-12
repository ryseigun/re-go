package basic

import "fmt"

func Map() {
	// How to make map 1
	book1 := map[string]string{
		"Title":  "Fahrenheit 451",
		"Author": "Ray Bradbury",
		"Year":   "1952",
	}

	// How to make map 2
	book2 := make(map[string]string)
	fmt.Println(book2)

	// Delete k,v
	delete(book1, "Year")
	fmt.Println(book1)
}
