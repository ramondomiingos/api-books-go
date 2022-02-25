package main

import "example/hello/server"

func main() {
	server := server.NewServer()
	server.Run()
}