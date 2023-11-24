package main

import (
	"fmt"
	"time"

	"root/modules/config"
	"root/modules/db"
	"root/modules/todo"

	"github.com/google/uuid"
)

func main() {
	config.Init()
	db.ConnectSqlite()

	id := uuid.New()

	fmt.Println("***** Welcome to TODO CLI *****")
	title, des := PipelineScan()

	newTodo := todo.Todo{
		ID:          id.String(),
		Title:       title,
		Description: des,
		IsCompleted: false,
		CreatedAt:   time.Now().String(),
	}
	todo.CreateTodo(newTodo)
}

func PipelineScan() (string, string) {
	title := ""
	des := ""

	fmt.Println("Enter your title: ")
	fmt.Print("> ")
	fmt.Scanln(&title)

	fmt.Println("Enter your description: ")
	fmt.Print("> ")
	fmt.Scanln(&des)

	return title, des
}
