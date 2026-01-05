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

// Embedding: you can assume it is like extending the class
type Admin struct {
	email    string
	password string
	User     // This line does the trick. here we are using anonymous decalration
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

func New(firstName string, lastName string, birthDate string, age int) (*User, error) {
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

// Contruction function for Admin
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: "Mango",
			LastName:  "Orange",
			BirthDate: "02/01/2025",
			Age:       3,
			createdAt: time.Now(),
		},
	}
}
