package todo

import (
	"fmt"
	"os"
)

func ClearAll(dbFile *os.File) {
	RemoveAllTodo(dbFile)
	fmt.Println("Clear Todo successfully")
}
