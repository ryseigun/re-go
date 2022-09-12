package basic

import (
	"fmt"
	"sort"
)

type User struct {
	name string
	age  int
}

type UserSlice []User

func (slc UserSlice) Len() int {
	return len(slc)
}

func (slc UserSlice) Less(i, j int) bool {
	return slc[i].age > slc[j].age
}

func (slc UserSlice) Swap(i, j int) {
	slc[i], slc[j] = slc[j], slc[i]
}

func SortThis() {
	users := []User{
		{"A", 9},
		{"B", 7},
		{"C", 10},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
