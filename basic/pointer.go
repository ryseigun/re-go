package basic

import "fmt"

type Address struct {
	City, Province, Country string
}

func Pointer() {
	// Passing by value
	ad1 := Address{"Solo", "Jawa Tengah", "Indonesia"}
	ad2 := ad1

	ad1.City = "Demak"
	fmt.Println(ad2) // add2 will not change according to add1

	// Passing by reference
	ad3 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	var ad4 *Address = &ad3

	ad3.City = "Tangerang"
	fmt.Println(ad4)

	// Change entire variable with same reference (wrong)
	ad5 := Address{"Jepara", "Jawa Tengah", "Indonesia"}
	ad6 := &ad5

	ad6 = &Address{"Medan", "Sumatera Utara", "Indonesia"}
	fmt.Println("ad5:", ad5) // ad5 will not change according to ad6
	fmt.Println("ad6", ad6)

	// Change entire variable with same reference (properly)
	ad7 := &ad5
	*ad7 = Address{"Semarang", "Jawa Tengah", "Indonesia"}
	fmt.Println("ad5:", ad5) // ad5 will change according to ad6
	fmt.Println("ad7", ad7)
}
