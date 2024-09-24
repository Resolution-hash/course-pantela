package handlers

import (
	"context"

	"github.com/Resolution-hash/course-pantela/task1/internal/services/messageService"
	"github.com/Resolution-hash/course-pantela/task1/internal/web/messages"
)

type messageHandler struct {
	Service *messageService.MessageService
}



// DeleteMessagesId implements messages.StrictServerInterface.
func (h *messageHandler) DeleteMessagesId(ctx context.Context, request messages.DeleteMessagesIdRequestObject) (messages.DeleteMessagesIdResponseObject, error) {
	messageID := request.Id

	err := h.Service.DeleteMessageByID(messageID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// PatchMessagesId implements messages.StrictServerInterface.
func (h *messageHandler) PatchMessagesId(ctx context.Context, request messages.PatchMessagesIdRequestObject) (messages.PatchMessagesIdResponseObject, error) {
	messageRequest := request.Body
	messageID := request.Id

	messageToUpdate := messageService.Message{Text: *messageRequest.Message}
	_, err := h.Service.UpdateMessageByID(messageID, messageToUpdate)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetMessages implements messages.StrictServerInterface.
func (h *messageHandler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	allMessages, err := h.Service.GetAllMessage()
	if err != nil {
		return nil, err
	}

	response := messages.GetMessages200JSONResponse{}

	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.ID,
			Message: &msg.Text,
		}
		response = append(response, message)
	}

	return response, nil
}

// PostMessages implements messages.StrictServerInterface.
func (h *messageHandler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	messageRequest := request.Body
	messageToCreate := messageService.Message{Text: *messageRequest.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)
	if err != nil {
		return nil, err
	}

	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}

	return response, nil
}

func NewMessageHandler(service *messageService.MessageService) *messageHandler {
	return &messageHandler{Service: service}
}

