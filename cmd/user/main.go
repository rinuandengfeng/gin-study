package main

import (
	"log"

	"gin-study/internal/user"
)

func main() {
	if err := user.Run("./config/user.yaml"); err != nil {
		log.Fatal(err)
	}
}
