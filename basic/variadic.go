package basic

import "fmt"

// How to make variadic function
func Sum(nums ...int) (total int) {
	total = 0
	for _, num := range nums {
		total += num
	}

	return
}

func Variadic() {
	fmt.Println(Sum())          // With no parameter
	fmt.Println(Sum(10, 10, 1)) // With int as parameter

	slice := []int{10, 9, 10}
	fmt.Println(Sum(slice...)) // With slice as parameter
}
