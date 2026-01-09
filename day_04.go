package main

import (
	"fmt"
	"time"
)

type Address struct {
	village      string
	street       string
	plotNumber   string
	city         string
	state        string
	pincode      int
	mobileNumber string
}
type Person struct {
	firstName string
	lastName  string
	age       int
	earning   float64
	address   *Address
	createdAt time.Time
}

func (person *Person) updatePerson(fName string, lName string, pAge int, pEarning float64) {
	if fName != "" {
		person.firstName = fName
	}
	if lName != "" {
		person.lastName = lName
	}
	if pAge > 0 {
		person.age = pAge
	}
	if pEarning > 100000 {
		person.earning = pEarning
	}
}

func (person Person) outputPersonDetails() {
	fmt.Println(person.firstName)
	fmt.Println(person.lastName)
	fmt.Println(person.age)
	fmt.Println(person.earning)
	fmt.Println(person.createdAt)
	fmt.Println(person.address.plotNumber)
	fmt.Println(person.address.street)
	fmt.Println(person.address.village)
	fmt.Println(person.address.city)
	fmt.Println(person.address.state)
	fmt.Println(person.address.pincode)
	fmt.Println(person.address.mobileNumber)
	fmt.Println("===================================")
}

// Creating constructors, Go uses prefers all constructor names should start with New[StructName]
// For nested thing we should always go for inner one first.

func NewAddress(empVill string, empCity string, empStreet string, empPlotNumber string, empState string, empPincode int, empMobileNumber string) *Address {
	return &Address{
		village:      empVill,
		city:         empCity,
		street:       empStreet,
		plotNumber:   empPlotNumber,
		state:        empState,
		pincode:      empPincode,
		mobileNumber: empMobileNumber,
	}
}

func NewPerson(empFirstName string, empLastName string, empAge int, empEarning float64, empAddress *Address) Person {
	return Person{
		firstName: empFirstName,
		lastName:  empLastName,
		age:       empAge,
		earning:   empEarning,
		address:   empAddress,
		createdAt: time.Now(),
	}
}

func day04Main() {
	empFirstName := "Shashi"
	empLastName := "Bhagat"
	empAge := 32
	empEarning := 1800000.00

	empVill := "Ghaghra"
	empStreet := "College Road"
	empCity := "Gumla"
	empPlotNumber := "376"
	empState := "Jharkhand"
	empPincode := 835208
	empMobileNumber := "9931920057"

	var empAddress *Address
	empAddress = &Address{
		village:      empVill,
		city:         empCity,
		street:       empStreet,
		plotNumber:   empPlotNumber,
		state:        empState,
		pincode:      empPincode,
		mobileNumber: empMobileNumber,
	}
	var employee Person
	employee = Person{
		firstName: empFirstName,
		lastName:  empLastName,
		age:       empAge,
		earning:   empEarning,
		address:   empAddress,
		createdAt: time.Now(),
	}

	fmt.Println(employee)

	employee.outputPersonDetails()
	employee.updatePerson("Bablu", "", -1, 1000.99)
	employee.outputPersonDetails()

}
