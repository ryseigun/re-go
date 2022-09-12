package basic

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func FunctionValue() {
	sayGoodBye := getGoodBye
	fmt.Println(sayGoodBye("Weakness"))
}
