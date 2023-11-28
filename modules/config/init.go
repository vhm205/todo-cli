package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	fmt.Println("Initializing config")

	func() {
		err := godotenv.Load()

		if err != nil {
			panic("Error loading .env file")
		}

		log.SetPrefix("LOG: ")
		log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	}()
}
