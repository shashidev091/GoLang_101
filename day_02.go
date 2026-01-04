package main

import "fmt"

func DayTwo() {
	fmt.Println("This is Day 2 of Go Bootcamp!")
	// loops()
	for i := 0; i <= 3; i++ {
		fmt.Println("Please select from the given options...")
		fmt.Println("Option 1: To see all the prime number included in the number of your choice.")
		fmt.Println("Option 2: To print the start pattern.")
		fmt.Print("Enter your option and press enter: ")
		var option int
		fmt.Scan(&option)
		switchCase(option)
	}
}

func loops() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}

func findPrimeNumbers(number int) {
	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			fmt.Print(i, ", ")
		} else {
			continue
		}
	}
}

func switchCase(choice int) {
	switch choice {
	case 1:
		fmt.Println("You choosed to find the Print numbers.")
		var askedNumber int
		fmt.Print("Please enter the number to get the list of prime numbers: ")
		fmt.Scan(&askedNumber)
		findPrimeNumbers(askedNumber)
		fmt.Println("================END==================")
	case 2:
		var patternSymbol string
		fmt.Println("You choosed to print the star pattern")
		fmt.Print("Choose your symbol to be used to generate start pattern. For example (*, #, emoji...):  ")
		fmt.Scan(&patternSymbol)
		for i := 0; i <= 5; i++ {
			for j := 0; j <= i; j++ {
				fmt.Print(patternSymbol)
			}
			fmt.Println()
		}
		fmt.Println("================END==================")
	default:
		fmt.Println("Wrong Input")
	}
}
