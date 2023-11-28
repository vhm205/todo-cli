package todo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

func CreateTodo(dbFile *os.File) {
	id := uuid.New()
	title, description := PipelineScan()

	newTodo := Todo{
		ID:          id.String(),
		Title:       title,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	var data []Todo = ReadFromFile(dbFile)

	if CheckExists(data, newTodo.Title) {
		log.Fatal("Todo already exists")
		return
	}

	data = append(data, newTodo)
	WriteToFile(data, dbFile)

	fmt.Println("Create Todo successfully")
}

func PipelineScan() (string, string) {
	var title string
	var description string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("---------- Create Todo ----------")
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
