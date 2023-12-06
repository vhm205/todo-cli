package todo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Update(dbFile *os.File) {
	unique, todo := UpdatePipelineScan(dbFile)

	todoCondition := Todo{ID: unique, Title: unique}
	isSuccess := UpdateTodo(todoCondition, todo, dbFile)

	if !isSuccess {
		fmt.Println("Todo not found")
		return
	}

	fmt.Println("Update Todo successfully")
}

func UpdatePipelineScan(dbFile *os.File) (string, Todo) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ID or Title: ")
	unique, _ := reader.ReadString('\n')
	unique = strings.TrimSpace(unique)

	todoCondition := Todo{ID: unique, Title: unique}
	todo, index := FindTodo(todoCondition, dbFile)

	if index == -1 {
		fmt.Println("Todo not found")
		return unique, todo
	}

	fmt.Println("ID:", todo.ID)
	fmt.Println("Title (2):", todo.Title)
	fmt.Println("Description (3):", todo.Description)
	fmt.Println("IsCompleted (4):", todo.IsCompleted)
	fmt.Println("CreatedAt:", todo.CreatedAt)
	fmt.Println("===================")

	fmt.Print("Select field for update (empty for skip): ")
	choice, _ := reader.ReadString('\n')

	switch strings.TrimSpace(choice) {
	case "2":
		fmt.Print("> ")
		newTitle, _ := reader.ReadString('\n')
		todo.Title = strings.TrimSpace(newTitle)
	case "3":
		fmt.Print("> ")
		newDescription, _ := reader.ReadString('\n')
		todo.Description = strings.TrimSpace(newDescription)
	case "4":
		todo.IsCompleted = !todo.IsCompleted
	}

	return unique, todo
}
