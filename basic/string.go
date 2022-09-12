package basic

import (
	"fmt"
	"strings"
)

func String() {
	fmt.Println(strings.Contains("Keiwa Sakurai", "ura"))
	fmt.Println(strings.Split("Keiwa Sakurai", " ")[0])
	fmt.Println(strings.ToLower("Keiwa Sakurai"))
	fmt.Println(strings.ToUpper("Keiwa Sakurai"))
	fmt.Println(strings.ToTitle("KEIWA SAKURAI-"))
	fmt.Println(strings.Trim("-Keiwa Sakurai--", "-"))
	fmt.Println(strings.ReplaceAll("Keiwa Sakurai", "a", "4"))
}
