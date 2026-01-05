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

	fmt.Println(appUser.firstName, appUser.lastName)
	printUserDetails(appUser)
	printUserDetailsPointers(&appUser)

	// Using attched methods
	appUser.outputUserDetails()

	// Manupulating struct variables
	appUser.clearUserName()
	appUser.outputUserDetails()

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

// Attaching a function to a struct
//
//	Receiver argument
func (user User) outputUserDetails() {
	fmt.Println("============Struct attached function =================")
	fmt.Println(user.firstName)
	fmt.Println(user.lastName)
	fmt.Println(user.birthDate)
	fmt.Println(user.age)
	fmt.Println(user.createdAt)
}

// To maupulate the variables in the struct you should always pass the pointer so that orignal struct is passed not the copy.
func (user *User) clearUserName() {
	fmt.Println("Clearing the firstName and Lastname ==================> ")
	user.firstName = ""
	user.lastName = ""
}

func printUserDetails(user User) {
	fmt.Println(user.firstName)
	fmt.Println(user.lastName)
	fmt.Println(user.birthDate)
	fmt.Println(user.age)
	fmt.Println(user.createdAt)
}

func printUserDetailsPointers(user *User) {
	// it works directly without derefencing
	fmt.Println(user.firstName)
	fmt.Println(user.lastName)
	fmt.Println(user.birthDate)
	fmt.Println(user.age)
	fmt.Println(user.createdAt)

	// although this is the way it should have if go did not have created the shortcut.
	// 	fmt.Println((*user).firstName)
	// 	fmt.Println((*user).lastName)
	// 	fmt.Println((*user).birthDate)
	// 	fmt.Println((*user).age)
	// 	fmt.Println((*user).createdAt)
}
