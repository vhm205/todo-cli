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
	fmt.Println("==========")

	switch choice {
	case 1:
		todo.Create(dbFile)
		break
	case 2:
		todo.GetList(dbFile)
		break
	case 3:
		todo.Update(dbFile)
		break
	case 4:
		todo.Delete(dbFile)
		break
	case 5:
		todo.ClearAll(dbFile)
		break
	default:
		os.Exit(0)
	}
}
