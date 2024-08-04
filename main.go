package main

import (
	"bufio"
	"os"

	"ukashanoor/noteapp/notes"
)

func main() {
	title, content := getInput()
	notes, err := notes.New(title, content)
	if err != nil {
		panic(1)
	}
	notes.PrintNote()
}

func getInput() (string, string) {
	var title string
	var content string

	reader := bufio.NewReader(os.Stdin)
	title, err := reader.ReadString('\n')
	if err != nil {
		panic(1)
	}
	content, err = reader.ReadString('\n')
	if err != nil {
		panic(1)
	}
	return title, content
}
