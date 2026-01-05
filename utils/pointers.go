package utils

import "fmt"

func pointerMain() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age: ", agePointer, *agePointer)

	adultYears := getAdultYears(*agePointer)
	fmt.Println(adultYears)

	editAgeToAdultYears(agePointer)
	fmt.Println("Adult age is:", age)
}

func getAdultYears(age int) int {
	return age - 18
}

func editAgeToAdultYears(age *int) {
	// Example of mutation
	if age == nil {
		age = new(int)
	}

	// we are changing the parent variable value in the memory directly.
	*age = *age - 18
}
