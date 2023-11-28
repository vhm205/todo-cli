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

func CheckExists(data []Todo, title string) bool {
	for _, item := range data {
		if item.Title == title {
			return true
		}
	}

	return false
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
