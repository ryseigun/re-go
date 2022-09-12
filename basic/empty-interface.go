package basic

import "fmt"

// Inteface kosong adalah interface yang tidak memiliki deklarasi method satupun
// secara otomatis semua tipe data akan menjadi implementasinya

// Fungsi yang dapat mengembalikan tipe data apa saja.
func Ups(i int) interface{} {
	if i == 1 {
		return 2
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func EmptyInterface() {
	var result interface{}
	result = Ups(1)

	fmt.Println(result)
}
