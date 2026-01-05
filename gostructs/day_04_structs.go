package gostructs

import (
	"fmt"
	"time"
)

func GoStructsMain() {
	var userFirstName string = getUserData("Please enter your first name: ")
	var userLastName string = getUserData("Please enter your last name: ")
	var userBirthdate string = getUserData("Please enter your birthdate (DD/MM/YY): ")

	fmt.Println(userFirstName, userLastName, userBirthdate)

	// Using Structs

	var appUser User
	appUser = User{
		firstName: userFirstName,
		lastName:  userLastName,
		age:       31,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	//or you can omit the struct keys and give the data in the same squence.

	// var appUser2 User
	// appUser2 = User{
	// 	31,
	// 	userBirthdate,
	// 	time.Now(),
	// 	userFirstName,
	// 	userLastName,
	// }

	// appUser2 = User{} => This will create a nil value of the object/struct
}

func getUserData(promptText string) string {
	var userInput string
	fmt.Print(promptText)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		panic("Invalid input")
	}
	return userInput
}

// Structs: are similar to class, which is used to group realated thing togeather

type User struct {
	age       int
	birthDate string
	createdAt time.Time
	firstName string
	lastName  string
}
