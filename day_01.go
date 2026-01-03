package main

import (
	"fmt"
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
}
