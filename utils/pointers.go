package utils

import "fmt"

func pointerMain() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age: ", agePointer, *agePointer)

	adultYears := getAdultYears(*agePointer)
	fmt.Println(adultYears)

	adultYearsWithPointers := getAdultYearsWithPointers(agePointer)
	fmt.Println("Adult age is:", adultYearsWithPointers)
}

func getAdultYears(age int) int {
	return age - 18
}

func getAdultYearsWithPointers(age *int) int {
	if age == nil {
		age = new(int)
	}
	return *age - 18
}
