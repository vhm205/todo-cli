package menu

import (
	"fmt"
	"os"
	"root/modules/todo"
)

func ShowMenu(dbFile *os.File) {
	fmt.Println("1. Create Todo")
	fmt.Println("2. List Todo")
	fmt.Println("3. Update Todo")
	fmt.Println("4. Delete Todo")
	fmt.Println("5. Clear Todo")
	fmt.Println("0. Exit")

	choice := int8(0)
	fmt.Print("Select an option: ")
	fmt.Scanln(&choice)
	fmt.Println(choice)

	switch choice {
	case 1:
		todo.CreateTodo(dbFile)
		break
	case 2:
		todo.ListTodo(dbFile)
		break
	case 5:
		todo.ClearTodo(dbFile)
	case 0:
		os.Exit(0)
	}
}
