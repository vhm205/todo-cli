package todo

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Delete(dbFile *os.File) {
	unique := DeletePipelineScan()
	todoCondition := Todo{ID: unique, Title: unique}
	todo, index := FindTodo(todoCondition, dbFile)

	if index == -1 {
		log.Println("Todo not found")
		return
	}

	DeleteTodo(todo, dbFile)
	fmt.Println("Delete Todo successfully")
}

func DeletePipelineScan() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ID or Title: ")
	unique, _ := reader.ReadString('\n')

	return unique
}
