package basic

import "fmt"

// Every variable has defaut value, ie: int with 0 and string with ""
// nil is and empty value or placeholder for empty map, interface, function, slice
// pointer, channel
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}

	return map[string]string{
		"name": name,
	}
}

func Nil() {
	// But this will printed as empty map, not nil.
	m := NewMap("")
	fmt.Println(m)
}
