package todo

import (
	"fmt"
	"os"
)

func ClearTodo(dbFile *os.File) {
	_ = os.WriteFile(dbFile.Name(), []byte{}, 0644)
	fmt.Println("Clear Todo successfully")
}
