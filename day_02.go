package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile string = "balance.txt"

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
		if option == 5 {
			break
		}
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

func drawTriangle(number int) {
	for i := 0; i <= number; i++ {
		for j := number; j <= 0; j-- {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func drawRightAngleTriangle() {
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
}

func drawRightAngleTriangleReverse(number int) {
	var patternSymbol string
	fmt.Println("You choosed to print the star pattern of reverse right angle triangle.")
	fmt.Print("Choose your symbol to be used to generate start pattern. For example (*, #, emoji...):  ")
	fmt.Scan(&patternSymbol)
	for i := 0; i <= number; i++ {
		for j := number; j >= i; j-- {
			fmt.Print(patternSymbol)
		}
		fmt.Println()
	}
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func readBalanceFromFile() (float64, error) {
	defaultValue := 10000.00
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return defaultValue, errors.New("Failed to read the file.")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return defaultValue, errors.New("Failed to parse the float value.")
	}

	return balance, nil
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
		drawRightAngleTriangle()
		fmt.Println("================END==================")
	case 3:
		drawRightAngleTriangleReverse(5)
		fmt.Println("================END==================")
	case 4:
		accountBalance, err := readBalanceFromFile()

		if err != nil {
			fmt.Println("ERROR==============>")
			fmt.Println(err)
			fmt.Println("===================>")
		}

		for {
			fmt.Println("What do you want to do?")
			fmt.Println("1. Check balance")
			fmt.Println("2. Deposit money")
			fmt.Println("3. Withdraw money")
			fmt.Println("4. Exit")

			var innerChoice int
			fmt.Print("Your choice: ")
			fmt.Scan(&innerChoice)

			switch innerChoice {
			case 1:
				fmt.Println("Your balance is", accountBalance)
			case 2:
				fmt.Println("Your deposit: ")
				var depositAmount float64
				fmt.Scan(&depositAmount)

				if depositAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}

				accountBalance += depositAmount
				fmt.Println("Balance updated! New amount: ", accountBalance)
				writeBalanceToFile(accountBalance)
			case 3:
				fmt.Print("Withdraw amount: ")
				var withdrawalAmount float64
				fmt.Scan(&withdrawalAmount)

				if withdrawalAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}

				if withdrawalAmount > accountBalance {
					fmt.Println("Invalid amount. You can't withdraw more than you have.")
					continue
				}

				accountBalance -= withdrawalAmount
				fmt.Println("Balance updated! New amount: ", accountBalance)
				writeBalanceToFile(accountBalance)
			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thanks for banking with us.")
				return
			}
		}
	default:
		fmt.Println("Goodbye!")
		fmt.Println("Thanks for banking with us.")
		return
	}
}
