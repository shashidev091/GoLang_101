package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("🔴 ➡️ Invalid Input")
	}
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

func (note Note) DisplayNote() {
	fmt.Printf("Your note titled %v has the following content: \n\n %v", note.title, note.content)

	fmt.Println(note.title)
	fmt.Println(note.content)
	fmt.Println(note.createdAt.Format(time.RFC822))
}
