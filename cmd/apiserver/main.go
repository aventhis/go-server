package main

import (
	"fmt"
	"github.com/aventhis/go-server/handlers"
	"net/http"
)

func main() {
	//  Запрашиваем случайный порт передав ":0"
	//listener, err := net.Listen("tcp", ":0")

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)
	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}
}
