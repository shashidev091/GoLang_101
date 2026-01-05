package main

import (
	"fmt"

	"github.com/shashidev091/GoLang_101/day_03/utils"
)

func day_three() {
	fmt.Println("This is Day 03, with new module.")
	const accountBalanceFile string = "balance.txt"
	balance, _ := utils.ReadFloatFromFile(accountBalanceFile)
	fmt.Println(balance)

	utils.UtilsMain()
}
