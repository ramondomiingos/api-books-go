package main

import (
	"api-books-go/database"
	"api-books-go/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
