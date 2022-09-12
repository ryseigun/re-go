package basic

import (
	"fmt"
	"strconv"
)

func StrConv() {
	// String -> Bool
	boolConv, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if boolConv {
			fmt.Println("Benar!")
		} else {
			fmt.Println("Salah!")
		}
	}

	// String -> int
	intConv, err := strconv.ParseInt("0000000091", 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(intConv)
	}
}
