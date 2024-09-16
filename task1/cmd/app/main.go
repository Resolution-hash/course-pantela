package main

import (
	"fmt"
	"net/http"

	database "github.com/Resolution-hash/course-pantela/task1/internal/database"
	"github.com/Resolution-hash/course-pantela/task1/internal/handlers"
	"github.com/Resolution-hash/course-pantela/task1/internal/messageService"
	"github.com/gorilla/mux"
)

// type Message struct {
// 	Text string `json:"text"`
// }

// type ErrorResponse struct {
// 	Code    int    `json:"code"`
// 	Message string `json:"message"`
// }

// func sendError(w http.ResponseWriter, code int, message string) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	errorResponse := ErrorResponse{
// 		code,
// 		message,
// 	}
// 	json.NewEncoder(w).Encode(errorResponse)
// }

// func GetHandler(w http.ResponseWriter, r *http.Request) {
// 	var messages []models.Message
// 	result := db.DB.Find(&messages)
// 	if result.Error != nil {
// 		sendError(w, http.StatusConflict, result.Error.Error())
// 		log.Println(result.Error.Error())
// 		return
// 	}

// 	if len(messages) != 0 {
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(messages)
// 		log.Println(messages)
// 		return
// 	}
// 	log.Println("Slice is empty")
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(Message{Text: "Slice is empty"})

// }

// func PostHandler(w http.ResponseWriter, r *http.Request) {
// 	data, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		sendError(w, http.StatusBadRequest, err.Error())
// 		log.Println(err)
// 		return
// 	}

// 	var message models.Message
// 	err = json.Unmarshal(data, &message)
// 	if err != nil {
// 		sendError(w, http.StatusBadRequest, err.Error())
// 		log.Println(err)
// 		return
// 	}

// 	result := db.DB.Create(&message)
// 	if result.Error != nil {
// 		sendError(w, http.StatusConflict, result.Error.Error())
// 		log.Println(result.Error.Error())
// 		return
// 	}
// 	log.Println("Message added successfully")
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(Message{Text: "Message added successfully"})

// }

// func DeleteHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	result := db.DB.Delete(models.Message{}, id)
// 	if result.Error != nil {
// 		sendError(w, http.StatusConflict, result.Error.Error())
// 		log.Println(result.Error.Error())
// 		return
// 	}
// 	log.Println("Message ", id, " deleted successfully")
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(Message{Text: "Message deleted successfully"})

// }

// func PatchHandler(w http.ResponseWriter, r *http.Request) {
// 	data, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		sendError(w, http.StatusBadRequest, err.Error())
// 		log.Println(err.Error())
// 		return
// 	}
// 	var message models.Message
// 	err = json.Unmarshal(data, &message)
// 	if err != nil {
// 		sendError(w, http.StatusBadRequest, err.Error())
// 		log.Println(err.Error())
// 		return
// 	}

// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	result := db.DB.Model(&models.Message{}).Where("id = ?", id).Update("text", message.Text)

// 	if result.Error != nil {
// 		sendError(w, http.StatusBadRequest, result.Error.Error())
// 		log.Println(result.Error.Error())
// 		return
// 	}

// 	log.Println("Updated: ", result.RowsAffected)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(Message{Text: "Data updated successfully"})

// }

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messageService.Message{})

	repo := messageService.NewMessageRepository(database.DB)
	service := messageService.NewMessageService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetMesssageHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/delete/{id}", handler.DeleteMessageHandler).Methods("DELETE")
	router.HandleFunc("/api/patch/{id}", handler.PatchMessageHandler).Methods("PATCH")

	fmt.Println("Server is starting on port:8080")
	http.ListenAndServe(":8080", router)
}
