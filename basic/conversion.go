package basic

import "fmt"

func Conversion() {
	// Same number but different size
	var nilai32 int32 = 300
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32, nilai64, nilai8)
	// How to count the overflow in int8?
	// max int8 = 128
	// min int8 = -127
	// 300 - 128 = -172
	// 172 + -127 = 45
	// 45 - 1 = 44

	// Character in string
	skill := "Destreza"
	var e byte = skill[0] // byte is tipealias of uint8
	var eString string = string(e)
	fmt.Println(skill, eString)
}
