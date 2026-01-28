package main

import "fmt"

func Day06Main() {
	fmt.Println("Day )> 6")
	someOrange()
}

func someOrange() {
	fmt.Println("Just a test. 🦥")
	printSomething("This is intersting")
	printSomething(1994)
	op := addGeneric(10, 20)
	fmt.Println(op)
}

// Generic in GoLang
func addGeneric[T int | float64](a, b T) T {
	return a + b
}
