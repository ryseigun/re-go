package basic

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name    string `required:"true"`
	Number  int
	IsValid bool
}

func Reflect() {
	sample1 := Sample{"A", 2, true}
	sample1Type := reflect.TypeOf(sample1)
	sample1Value := reflect.ValueOf(sample1)

	for i := 0; i < sample1Type.NumField(); i++ {
		x := sample1Type.Field(i)
		fmt.Println(x.Name)                // Field name
		fmt.Println(x.Type)                // Field data type
		fmt.Println(x.Tag.Get("required")) // if tag is not exist, it will return empty
		fmt.Println(sample1Value.Field(i)) // Value of the field
	}
}
