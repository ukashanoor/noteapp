package notes

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content should not be empty")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note *Note) PrintNote() {
	fmt.Printf("Note Title: %v, Content: %v", note.Title, note.Content)
}
