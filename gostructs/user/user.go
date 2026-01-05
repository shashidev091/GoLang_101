package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	Age       int
	BirthDate string
	createdAt time.Time
	FirstName string
	LastName  string
}

// Attaching a function to a struct
//
//	Receiver argument
func (user User) OutputUserDetails() {
	fmt.Println("============Struct attached function =================")
	fmt.Println(user.FirstName)
	fmt.Println(user.LastName)
	fmt.Println(user.BirthDate)
	fmt.Println(user.Age)
	fmt.Println(user.createdAt)
}

// To maupulate the variables in the struct you should always pass the pointer so that orignal struct is passed not the copy.
func (user *User) ClearUserName() {
	fmt.Println("Clearing the firstName and Lastname ==================> ")
	user.FirstName = ""
	user.LastName = ""
}

func NewUser(firstName string, lastName string, birthDate string, age int) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		BirthDate: birthDate,
		createdAt: time.Now(),
	}
}

func NewUserV2(firstName string, lastName string, birthDate string, age int) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, Last name and birthdate are required parameters.")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		BirthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
