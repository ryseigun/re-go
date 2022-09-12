package basic

import "fmt"

// Panic
func runApp(isError bool) {
	defer catchError()

	fmt.Println("Running the app...")
	if isError {
		panic("There's an error")
	}
	fmt.Println("Closing the app...")
}

// Recover
func catchError() {
	errMsg := recover()
	fmt.Println("Error message : ", errMsg)
}

// Main
func DeferRecover() {
	runApp(true)
}
