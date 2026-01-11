package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("🔴 ➡️ Invalid Input")
	}
	return Todo{Text: content}, nil
}

func (todo Todo) DisplayTodo() {
	fmt.Printf("Your todo has the following content: \n\n %v", todo.Text)
}

func (todo Todo) SaveFile() error {
	fileName := "todo.json"
	json_text, err := json.MarshalIndent(todo, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json_text, 0644)
}
