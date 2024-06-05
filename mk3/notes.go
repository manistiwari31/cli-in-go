package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title       string
	Description string
	CreatedAt   time.Time
}

var notes []Note

func main() {
	if len(os.Args) < 2 {
		listNotes()
		os.Exit(0)
	}

	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: add <title> <description>")
			os.Exit(1)
		}
		title := os.Args[2]
		description := os.Args[3]
		addNote(title, description)
	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: search <title>")
			os.Exit(1)
		}
		searchTitle := os.Args[2]
		searchByTitle(searchTitle)
	default:
		fmt.Println("Invalid command.")
		os.Exit(1)
	}
}

func addNote(title, description string) {
	note := Note{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}

	notes = append(notes, note)

	fmt.Println("Note added successfully!")
}

func listNotes() {
	fmt.Println("List of Notes:")
	if len(notes) == 0 {
		fmt.Println("No notes found.")
		return
	}

	for i, note := range notes {
		fmt.Printf("\nNote %d:\n", i+1)
		fmt.Printf("Title: %s\n", note.Title)
		fmt.Printf("Description: %s\n", note.Description)
		fmt.Printf("Created At: %s\n", note.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

func searchByTitle(searchTitle string) {
	var foundNotes []Note
	for _, note := range notes {
		if strings.Contains(strings.ToLower(note.Title), strings.ToLower(searchTitle)) {
			foundNotes = append(foundNotes, note)
		}
	}

	if len(foundNotes) == 0 {
		fmt.Println("No matching notes found.")
		return
	}

	fmt.Println("Matching Notes:")
	for i, note := range foundNotes {
		fmt.Printf("\nNote %d:\n", i+1)
		fmt.Printf("Title: %s\n", note.Title)
		fmt.Printf("Description: %s\n", note.Description)
		fmt.Printf("Created At: %s\n", note.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}
