package basic

import (
	"fmt"
)

func deleteEl(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func Array() {
	// How to make array directly
	arr := [3]int{
		1, 2, 3,
	}
	fmt.Println(arr)

	// 'Length' of the array
	var arr1 [4]int        // Walaupun belum diisi data
	fmt.Println(len(arr1)) // Panjangnya tetap 4

	// With slice
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var sliceMonths = months[4:7]
	fmt.Println(
		len(sliceMonths), // Panjang: 4,5,6 ada 3
		cap(sliceMonths)) // Capacity: 4 s/d len(months)

	// Appending item
	sliceMonths = append(sliceMonths, "randomStr")
	fmt.Println(sliceMonths)

	// Deleting item with index i
	sliceMonths = deleteEl(sliceMonths, 8)

	// Making string slice from scratch with len 2 and cap 5
	newSlice := make([]string, 2, 5)
	copy(newSlice, sliceMonths) // Dikopi sebanyak length (2 item)
	fmt.Println(newSlice)

	// Declaration difference array slice
	arrVar := [3]int{1, 2, 3} // or [...]int{}
	slcVar := []int{1, 2, 3}
	fmt.Println(arrVar, slcVar)

	// Make a exact copy from existing slice
	cpSlice := make([]string, len(sliceMonths), cap(sliceMonths))
	copy(cpSlice, sliceMonths)
	fmt.Println(cpSlice)
}
