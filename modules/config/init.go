package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Init() {
	fmt.Println("Init config...")

	func() {
		err := godotenv.Load()

		if err != nil {
			panic("Error loading .env file")
		}
	}()
}
