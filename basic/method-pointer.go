package basic

import "fmt"

type Man struct {
	name string
}

// Method without pointer
func (m Man) Married1() {
	m.name = "Mr. " + m.name
}

// Method with pointer
func (m *Man) Married2() {
	m.name = "Mr. " + m.name
}

func MethodPointer() {
	man1 := Man{"Shinnosuke"}
	man1.Married1()
	fmt.Println(man1)
	man1.Married2()
	fmt.Println(man1)
}
