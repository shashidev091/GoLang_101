package utils

import "fmt"

func pointerMain() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age: ", agePointer, *agePointer)

	adultYears := getAdultYears(*agePointer)
	fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
