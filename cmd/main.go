package main

import (
	"log"

	"github.com/max99xam/todo-app"
	"github.com/max99xam/todo-app/pkg/handler"
)

func main() {

	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}