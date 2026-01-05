package main

import (
	"fmt"

	"github.com/shashidev091/GoLang_101/day_03/gostructs"
	"github.com/shashidev091/GoLang_101/day_03/gostructs/user"
)

func day_three() {
	fmt.Println("This is Day 03, with new module.")
	// const accountBalanceFile string = "balance.txt"
	// balance, _ := utils.ReadFloatFromFile(accountBalanceFile)
	// fmt.Println(balance)

	// utils.UtilsMain()
	gostructs.GoStructsMain()
	var appUser *user.User
	appUser, err := user.NewUserV2("Shashi", "Bhagat", "20/04/1994", 31)

	if err != nil {
		fmt.Println(err)
	}

	appUser.OutputUserDetails()

}
