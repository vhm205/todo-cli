package todo

import (
	"fmt"
	"os"
)

func GetList(dbFile *os.File) []Todo {
	data := FindAllTodo(dbFile)
	for k, value := range data {
		k = k + 1
		fmt.Println("ID:", value.ID)
		fmt.Println("Title:", value.Title)
		fmt.Println("Description:", value.Description)
		fmt.Println("IsCompleted:", value.IsCompleted)
		fmt.Println("CreatedAt:", value.CreatedAt)
		fmt.Println("=======")
	}
	return data
}
