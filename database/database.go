package database

import "fmt"

var conn string

// This function will be executed if the package is called
func init() {
	fmt.Println("Init function called")
	conn = "MySQL"
}

func GetDatabase() string {
	return conn
}
