package main

import (
	"fmt"

	"github.com/shashidev091/GoLang_101/day_04/note"
)

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

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.DisplayNote()
	err = userNote.SaveFile()
	if err != nil {
		fmt.Println("Saving the note failed with the error.")
	}

	fmt.Println("File Saved successfully.")
}
