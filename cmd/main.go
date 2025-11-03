package main

import (
	"log"
	"todo"
	handler "todo/pkg/handlers"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("ОШИБКА ЗАПУСКА ПРИЛОЖЕНИЯ %s", err.Error())
	}

}
