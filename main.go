package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/suvam720/book-store/pkg/database"
	"github.com/suvam720/book-store/pkg/routers"
)

func main() {
	godotenv.Load(".env")
	EnvKey, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("failed to read env file")
	}
	dsn := EnvKey["DB_STRING"]
	err1 := database.Connect(dsn)
	if err1 != nil {
		log.Fatal(err)
	}
	routes := routers.Routes()
	routes.Run()
}
