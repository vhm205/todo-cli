package todo

import "fmt"

type Todo struct {
	ID          string
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   string
}

func CreateTodo(todo Todo) {
	fmt.Println("Create Todo successfully")
	fmt.Println(todo)
}
