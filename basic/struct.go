package basic

import "fmt"

// Define the struct
type KamenRider struct {
	Name, Driver string
	Year         int
}

// Struct method
func (kr KamenRider) RiderKick() {
	fmt.Println(kr.Name + " does a rider kick!")
}

func Struct() {
	// Make the struct 1
	var build KamenRider
	build.Name = "Kamen Rider Build"
	build.Driver = "Build Driver"
	build.Year = 2017
	fmt.Println(build)

	// Make the struct 2
	revice := KamenRider{
		Name:   "Kamen Rider Revice",
		Driver: "Revice Driver",
		Year:   2021,
	}
	fmt.Println(revice)

	// Make the struct 3 -> Rentan error
	double := KamenRider{"Kamen Rider W", "Double Driver", 2010}
	fmt.Println(double)

	// Call the struct method
	build.RiderKick()
	revice.RiderKick()
	double.RiderKick()
}
