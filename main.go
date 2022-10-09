package main

import (
	"fiber/fiber-graphql/config/server"
	"log"
)

const defaultPort = "8080"

func main() {
	app := server.InitServer()

	err := app.Listen(":" + defaultPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
