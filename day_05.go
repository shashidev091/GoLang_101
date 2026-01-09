package main

import "fmt"

//Types

type str string

func (text str) log() {
	fmt.Println(text)
}

func Day05Main() {
	var name str = "This world is nice."
	// fmt.Println(name)

	// TO use the explicit type we have to declare and use the explicit type
	name.log()
}
