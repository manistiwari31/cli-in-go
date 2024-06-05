package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Todo struct {
	idx           string
	prior         string
	body          string
	status        bool
	dateadded     string
	timeadded     time.Time
	timecompleted time.Time
}

var TodoList []*Todo // Changed to store pointers to Todo

func main() {
	var query string
	var allowquery bool
	if len(os.Args) >= 2 {
		query = os.Args[1]
		allowquery = true
	} else {
		for _, todo := range TodoList {
			color.Green(todo.body)
		}
	}

	if allowquery {
		query = strings.ToLower(query) // Convert query to lowercase
		if query == "add" {
			if len(os.Args) >= 4 {
				body := os.Args[2]
				prior := os.Args[3]
				t := newTodo(body, prior)
				TodoList = append(TodoList, t) // Append the pointer itself

				for _, todo := range TodoList {
					color.Green(todo.body) // Change todo.done to todo.status
				}
			}
		}

		if query == "show" {
			for _, todo := range TodoList {
				color.Green(todo.body) // Change todo.done to todo.status
			}
		}
		if query == "done" {
			if len(os.Args) >= 4 {
				index := os.Args[2]
				indexint, _ := strconv.Atoi(index)
				t := pick(indexint)
				t.done()
			}
		}
		if query == "undo" {
			if len(os.Args) >= 4 {
				index := os.Args[2]
				indexint, _ := strconv.Atoi(index)
				t := pick(indexint)
				t.undone()
			}
		}

		if query == "list" {
			if len(os.Args) >= 4 {
				args := os.Args[2]
				var doneornot bool
				if args == "done" {
					doneornot = true
				}
				for _, todo := range TodoList {
					if todo.status == doneornot {
						color.Blue(todo.body)
					}
				}
			}
		}
	}
}

func (t *Todo) done() {
	t.status = true
	t.timecompleted = time.Now()
}

func (t *Todo) undone() {
	t.status = false
	t.timecompleted = time.Time{} // nil
}

func pick(i int) *Todo {
	if i >= 0 && i < len(TodoList) {
		return TodoList[i] // Return the pointer itself
	}
	return nil
}

func newTodo(body, prior string) *Todo {
	color.Blue(body)

	return &Todo{
		body:      body,
		prior:     prior,
		timeadded: time.Now(),
		dateadded: fmt.Sprint(time.Now().Month(), time.Now().Weekday()),
	}
}
