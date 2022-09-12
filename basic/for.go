package basic

import "fmt"

func For() {
	// How to make for loop 1
	counter := 1
	for counter <= 5 {
		fmt.Println(counter)
		counter++
	}

	// How to make for loop 2
	// With init statement; condition; post statement
	for i := 0; i < 5; i++ {
		fmt.Println(i + 1)
	}

	// How to make for loop 3
	riders := []string{"Double", "Accel", "Revice", "Demons"}
	for _, e := range riders {
		fmt.Println(e)
	}

	// How to make for loop 4
	ultimateForms := map[string]string{
		"Kabuto": "Hyper Form",
		"Double": "Xtreme",
		"Build":  "Genius Form",
	}
	for k, v := range ultimateForms {
		fmt.Println("Kamen rider ", k, " = ", v)
	}
}
