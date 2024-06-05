package todo

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type Todo struct {
	prior         string
	body          string
	status        bool
	dateadded     string
	timeadded     time.Time
	timecompleted time.Time
}

var TodoList []Todo

func main() {
	var query string
	var allowquery bool
	if len(os.Args) >= 2 {
		query = os.Args[1]
		allowquery = true
	}

	if allowquery {
		if query == "add" || query == "Add" {

			if len(os.Args) >= 4 {
				body := os.Args[2]
				prior := os.Args[3]
				t := newTodo(body, prior)
				TodoList = append(TodoList, &t)
			}
			if query == "show" || query == "Show" {
				for _, todo := range TodoList {
					color.Green(todo.body, "/t", todo.done)
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

func newTodo(body, prior string) *Todo {
	color.Blue(body)

	return &Todo{
		// idx: c,
		body:      body,
		prior:     prior,
		timeadded: time.Now(),
		dateadded: fmt.Sprint(time.Now().Month(), time.Now().Weekday()),
	}

}
