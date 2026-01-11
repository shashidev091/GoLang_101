package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"` //Struck tags: also used for metadata
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("🔴 ➡️ Invalid Input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) DisplayNote() {
	fmt.Printf("Your note titled %v has the following content: \n\n %v", note.Title, note.Content)

	fmt.Println(note.Title)
	fmt.Println(note.Content)
	fmt.Println(note.CreatedAt.Format(time.RFC822))
}

func (note Note) SaveFile() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json_text, err := json.MarshalIndent(note, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json_text, 0644)
}
