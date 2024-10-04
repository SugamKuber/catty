package main

import (
	"fmt"
	"github.com/joho/godotenv"

	"server/internal/db"
	"server/internal/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db := database.NewDatabase()
	r := routes.SetupRouter(&db)
	fmt.Println("Server started on port 8080")
	r.Run(":8080")
}
