package basic

import "fmt"

func Alias() {
	// Creating alias, NoKTP is string
	type NoKTP string

	var noKTP1 NoKTP = "123456"
	fmt.Println(noKTP1)
}
