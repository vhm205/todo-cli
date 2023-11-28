package main

import (
	"fmt"

	"root/modules/config"
	"root/modules/db"
	"root/modules/menu"
)

func main() {
	config.Init()
	dbFile := db.OpenDbFile()

	fmt.Println("***** Welcome to TODO CLI *****")
	menu.ShowMenu(dbFile)

	defer dbFile.Close()
}
