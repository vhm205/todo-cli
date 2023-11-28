package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func OpenDbFile() *os.File {
	dbFolderName := "data"
	dbFileName := filepath.Join(dbFolderName, "database.json")

	_, errFolder := os.Stat(dbFolderName)
	dbFile, errFile := os.OpenFile(dbFileName, os.O_APPEND, 0755)

	if os.IsNotExist(errFolder) {
		err := os.MkdirAll(dbFolderName, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	if os.IsNotExist(errFile) {
		_, err := os.Create(dbFileName)

		if err != nil {
			log.Fatal(errFile)
		}
	}

	fmt.Println("Connect to database successfully!!")
	fmt.Println("------------------------------")
	return dbFile
}
