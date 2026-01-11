package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func getTodoData() string {
	return getUserInput("Todo text: ")
}

func saveData(data Saver) error {
	err := data.SaveFile()
	if err != nil {
		fmt.Println("Saving the note failed with the error.")
	}
	return nil
}
