package handlers

import (
	"context"

	"github.com/Resolution-hash/course-pantela/task1/internal/services/userService"
	"github.com/Resolution-hash/course-pantela/task1/internal/web/users"
)

type userHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *userHandler {
	return &userHandler{Service: service}
}

// DeleteUsersId implements users.StrictServerInterface.
func (h *userHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID := request.Id

	err := h.Service.DeleteUserByID(userID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetUsers implements users.StrictServerInterface.
func (h *userHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		formattedCreatedAt := usr.Model.CreatedAt.Format("2006-01-02 15:04:05")
		formattedUpdatedAt := usr.Model.UpdatedAt.Format("2006-01-02 15:04:05")
		// formattedDeletedAt := usr.Model.DeletedAt.Format("2006-01-02 15:04:05")
		user := users.Users{
			Id:        &usr.ID,
			Email:     &usr.Email,
			Password:  &usr.Password,
			CreatedAt: &formattedCreatedAt,
			UpdatedAt: &formattedUpdatedAt,
			// DeleteAt:  &formattedDeletedAt,
		}
		response = append(response, user)
	}
	return response, nil
}

// PatchUsersId implements users.StrictServerInterface.
func (h *userHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userRequest := request.Body
	userID := request.Id

	userToUpdate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	err := h.Service.PatchUserByID(userID, userToUpdate)
	if err != nil {
		return nil, err
	}

	return nil, err
}

// PostUsers implements messages.StrictServerInterface.
func (h *userHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	err := h.Service.PostUser(userToCreate)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
