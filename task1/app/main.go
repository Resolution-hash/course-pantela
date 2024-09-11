package main

import (
	"fmt"
	"io"
	"net/http"

	"encoding/json"

	"github.com/Resolution-hash/course-pantela/task1/db"
	"github.com/Resolution-hash/course-pantela/task1/models"
	"github.com/gorilla/mux"
)

type Message struct {
	Text string `json:"text"`
}

var message Message

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var messages []models.Message
	result := db.DB.Find(&messages)
	if result.Error != nil {
		fmt.Fprintln(w, result.Error)
		return
	}

	if len(messages) != 0 {
		fmt.Fprintln(w, messages)
		return
	}

	fmt.Fprintln(w, "Slice is empty")

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	err = json.Unmarshal(data, &message)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	result := db.DB.Create(&models.Message{
		Text: message.Text,
	})
	if result.Error != nil {
		fmt.Fprintln(w, result.Error)
		return
	}
	fmt.Fprintln(w, "Message added successfully")

}

func main() {
	db.InitDB()

	db.DB.AutoMigrate(&models.Message{})

	router := mux.NewRouter()

	router.HandleFunc("/api/getMessages", GetHandler).Methods("GET")
	router.HandleFunc("/api/postMessage", PostHandler).Methods("POST")

	fmt.Println("Server is starting on port:8080")
	http.ListenAndServe(":8080", router)
}
