package basic

import (
	"errors"
	"fmt"
)

// It's common in golang to return a result and  an error in a function
func dividexy(x int, y int) (r int, e error) {
	if y == 0 {
		return 0, errors.New("y musn't be 0")
	}

	return x / y, nil
}

func ErrorInterface() {
	fmt.Println(dividexy(1, 2))
	fmt.Println(dividexy(1, 0))
}
