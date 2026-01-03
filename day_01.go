package main // Used to group the files similar to java

import (
	"fmt" // core package, knows as format
	"math"
)

func main() {
	fmt.Println("Hello, Shashi 🌚")
	// var investmentAmount int = 1000
	// var expectedReturnRate float64 = 5.5
	// var years int = 10
	// var futureValue float64 = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	const inflationRate float64 = 2.4
	var investmentAmount, years int = 1000, 10
	expectedReturnRate := 5.5
	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println(futureValue)
	fmt.Printf("The future value is %.2f.\n", futureValue)

	var futureRealValue float64 = futureValue / math.Pow(1+inflationRate/100, float64(years))
	fmt.Println(futureRealValue)

	// Check your date of birth (Using only single words/digits with Scan)
	age := calculate_age()

	//Create a profit calulator
	// profitCalculator()

	var introduction = fmt.Sprintf("Hello, My name is Shashi Bhagat and I am %d years old. I am learning Go Language.", age)
	fmt.Println(introduction)
}

func calculate_age() int {
	const currentYear int = 2025
	var yearOfBirth uint
	var age int

	fmt.Print("Enter your year of birth: ")
	fmt.Scan(&yearOfBirth)
	age = currentYear - int(yearOfBirth)
	fmt.Printf("%T\n", currentYear)
	fmt.Println("You are", age, "Years old.")
	return age
}

func profitCalculator() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Please enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Please enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Please enter the Tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)

	ratio := ebt / profit

	// fmt.Println(ebt)
	// fmt.Println(profit)
	// fmt.Println(ratio)

	fmt.Printf("Your ebt is: %v.\nYour profit is: %v.\nYour ratio is : %.2f.", ebt, profit, ratio)
}
