package basic

import "fmt"

func divide(x, y int) int {
	return x / y
}

// Generate slice but one of the element
// is 0
func GenerateSlice() []int {
	var slice []int
	for i := 2; i <= 10; i += 2 {
		if i == 6 {
			slice = append(slice, 0)
		} else {
			slice = append(slice, i)
		}
	}

	return slice
}

// When panic will occur
func PanicRecover() {
	// Make a recover with defer
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error: ", err)
		}
	}()

	slice := GenerateSlice()

	for _, e := range slice {
		fmt.Println(e, " / 2 = ", divide(100, e))
	}
}
