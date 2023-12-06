package todo

import (
	"encoding/json"
	"log"
	"os"
)

type Todo struct {
	ID          string `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description"`
	IsCompleted bool   `json: "isCompleted"`
	CreatedAt   string `json: "createdAt"`
}

func InsertTodo(todo Todo, dbFile *os.File) {
	data := FindAllTodo(dbFile)
	data = append(data, todo)
	WriteToFile(data, dbFile)
}

func UpdateTodo(condition Todo, todo Todo, dbFile *os.File) bool {
	var data []Todo = FindAllTodo(dbFile)
	_, index := FindTodo(condition, dbFile)

	if index == -1 {
		return false
	}

	data[index] = todo
	WriteToFile(data, dbFile)
	return true

}

func FindAllTodo(dbFile *os.File) []Todo {
	return ReadFromFile(dbFile)
}

func FindTodo(todo Todo, dbFile *os.File) (Todo, int) {
	data := FindAllTodo(dbFile)

	for index, item := range data {
		if item.Title == todo.Title || item.ID == todo.ID {
			return item, index
		}
	}

	return Todo{}, -1
}

func DeleteTodo(todo Todo, dbFile *os.File) bool {
	_, index := FindTodo(todo, dbFile)

	if index == -1 {
		return false
	}

	data := FindAllTodo(dbFile)
	data = append(data[:index], data[index+1:]...)

	WriteToFile(data, dbFile)
	return true
}

func RemoveAllTodo(dbFile *os.File) {
	_ = os.WriteFile(dbFile.Name(), []byte{}, 0644)
}

func ReadFromFile(dbFile *os.File) []Todo {
	file, err := os.ReadFile(dbFile.Name())
	if err != nil {
		log.Fatal("Error reading file")
	}

	data := []Todo{}
	_ = json.Unmarshal(file, &data)

	return data
}

func WriteToFile(data []Todo, dbFile *os.File) {
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile(dbFile.Name(), file, 0644)
}
