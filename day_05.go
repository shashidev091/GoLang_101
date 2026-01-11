package main

import (
	"fmt"

	"github.com/shashidev091/GoLang_101/day_04/note"
	"github.com/shashidev091/GoLang_101/day_04/todo"
)

// interfaces: always give name that ends with "er", this is a design pattern.
type Saver interface {
	SaveFile() error
}

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

	todoText := getTodoData()
	todo, err := todo.New(todoText)

	if err != nil {
		// fmt.Println(err)
		panic(err)
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.DisplayTodo()

	saveData(todo)

	userNote.DisplayNote()
	saveData(userNote)

	// Interfaces
}
