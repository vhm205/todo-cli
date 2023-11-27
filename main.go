package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"root/modules/config"
	"root/modules/db"
	"root/modules/todo"

	"github.com/google/uuid"
)

func main() {
	config.Init()
	dbFile := db.OpenDbFile()

	id := uuid.New()

	fmt.Println("***** Welcome to TODO CLI *****")
	title, description := PipelineScan()

	newTodo := todo.Todo{
		ID:          id.String(),
		Title:       title,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now().String(),
	}
	todo.CreateTodo(newTodo, dbFile)

	defer dbFile.Close()
}

func PipelineScan() (string, string) {
	var title string
	var description string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your title: ")
	fmt.Print("> ")
	title, _ = reader.ReadString('\n')

	fmt.Println("Enter your description: ")
	fmt.Print("> ")
	description, _ = reader.ReadString('\n')

	title = strings.Replace(title, "\n", "", -1)
	description = strings.Replace(description, "\n", "", -1)

	return title, description
}
