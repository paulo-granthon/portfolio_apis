package schemas

import "models"

type GetUserRequest struct {
	Id uint64 `json:"id"`
}

type RegisterUserRequest struct {
	Payload models.RegisterUser `json:"payload"`
}

type RegisterUserResponse struct {
	Id uint64 `json:"id"`
}

type CreateUserRequest struct {
	User models.CreateUser `json:"user"`
}

type CreateUserResponse struct {
	Id uint64 `json:"id"`
}

type UpdateUserRequest struct {
	User models.User `json:"user"`
}
