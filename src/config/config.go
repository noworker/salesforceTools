package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitApp() {
	if os.Getenv("MODE") == "DEV" {
		err := godotenv.Load("../dev.env")
		if err != nil {
			panic("ERROR: do not load env file")
		}
		fmt.Println("SUCCESS: load env file")
	}
}
