package main

import (
	"log"

	database "github.com/Resolution-hash/course-pantela/task1/internal/database"
	messageHandlers "github.com/Resolution-hash/course-pantela/task1/internal/handlers/messages"
	userHandlers "github.com/Resolution-hash/course-pantela/task1/internal/handlers/users"
	"github.com/Resolution-hash/course-pantela/task1/internal/services/messageService"
	"github.com/Resolution-hash/course-pantela/task1/internal/services/userService"
	"github.com/Resolution-hash/course-pantela/task1/internal/web/messages"
	"github.com/Resolution-hash/course-pantela/task1/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&messageService.Message{})
	if err != nil {
		log.Fatal(err.Error())
	}
	messageRepo := messageService.NewMessageRepository(database.DB)
	messageService := messageService.NewMessageService(messageRepo)
	messageHandler := messageHandlers.NewMessageHandler(messageService)

	err = database.DB.AutoMigrate(&userService.User{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := userHandlers.NewUserHandler(userService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictMessageHandler := messages.NewStrictHandler(messageHandler, nil)
	messages.RegisterHandlers(e, strictMessageHandler)

	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
