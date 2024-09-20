package main

import (
	"fmt"
	"log"
	"net/http"

	database "github.com/Resolution-hash/course-pantela/task1/internal/database"
	"github.com/Resolution-hash/course-pantela/task1/internal/handlers"
	"github.com/Resolution-hash/course-pantela/task1/internal/messageService"
	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&messageService.Message{})
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := messageService.NewMessageRepository(database.DB)
	service := messageService.NewMessageService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetMesssageHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/delete/{id}", handler.DeleteMessageHandler).Methods("DELETE")
	router.HandleFunc("/api/patch/{id}", handler.PatchMessageHandler).Methods("PATCH")

	fmt.Println("Server is starting on port:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
