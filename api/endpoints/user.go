package endpoints

import (
	"models"
	"net/http"
	"schemas"
	"server"
	"strconv"
)

func UserEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{
			Path: "/register",
			Methods: []server.Method{
				server.NewMethod("POST", RegisterUser),
			},
		},
		{
			Path: "/users",
			Methods: []server.Method{
				server.NewMethod("GET", GetUsers),
				server.NewMethod("POST", CreateUser),
			},
		},
		{
			Path: "/users/{id}",
			Methods: []server.Method{
				server.NewMethod("GET", GetUser),
				server.NewMethod("PUT", UpdateUser),
				server.NewMethod("DELETE", DeleteUser),
			},
		},
	}
}

func GetUsers(s server.Server, w http.ResponseWriter, r *http.Request) error {
	userModule, err := s.Storage.GetUserModule()
	if err != nil {
		return err
	}

	users, err := userModule.Get()
	if err != nil {
		return server.SendError(w, err)
	}
	return server.WriteJSON(w, http.StatusOK, users)
}

func GetUser(s server.Server, w http.ResponseWriter, r *http.Request) error {
	idStr, err := server.GetRequestParam(r, "id")
	if err != nil {
		return server.SendError(w, err)
	}

	id, err := strconv.ParseUint(*idStr, 10, 64)
	if err != nil {
		return server.SendError(w, err)
	}

	userModule, err := s.Storage.GetUserModule()
	if err != nil {
		return err
	}

	user, err := userModule.GetById(id)
	if err != nil {
		return server.WriteJSON(w, http.StatusNotFound, server.Error{Error: "user not found"})
	}

	return server.WriteJSON(w, http.StatusOK, user)
}

func RegisterUser(s server.Server, w http.ResponseWriter, r *http.Request) error {
	var request schemas.RegisterUserRequest
	if error := server.ReadJSON(r, &request.Payload); error != nil {
		return server.SendError(w, error)
	}

	user := models.NewRegisterUser(
		request.Payload.Name,
		request.Payload.Password,
	)

	userModule, err := s.Storage.GetUserModule()
	if err != nil {
		return err
	}

	id, err := userModule.Register(user)
	if err != nil {
		return server.SendError(w, err)
	}

	return server.WriteJSON(w, http.StatusCreated, schemas.RegisterUserResponse{Id: *id})
}

func CreateUser(s server.Server, w http.ResponseWriter, r *http.Request) error {
	var request schemas.CreateUserRequest
	if error := server.ReadJSON(r, &request.User); error != nil {
		return server.SendError(w, error)
	}

	user := models.NewCreateUser(
		request.User.Name,
		request.User.Password,
		request.User.Summary,
		request.User.SemesterMatriculed,
		request.User.GithubUsername,
	)

	userModule, err := s.Storage.GetUserModule()
	if err != nil {
		return err
	}

	id, err := userModule.Create(user)
	if err != nil {
		return server.SendError(w, err)
	}

	return server.WriteJSON(w, http.StatusCreated, schemas.CreateUserResponse{Id: *id})
}

func UpdateUser(s server.Server, w http.ResponseWriter, r *http.Request) error {
	idStr, err := server.GetRequestParam(r, "id")
	if err != nil {
		return server.SendError(w, err)
	}

	id, err := strconv.ParseUint(*idStr, 10, 64)
	if err != nil {
		return server.SendError(w, err)
	}

	var request schemas.UpdateUserRequest
	if err := server.ReadJSON(r, &request); err != nil {
		return server.SendError(w, err)
	}

	user := models.NewUpdateUser(
		id,
		request.User.Name,
		request.User.Summary,
		request.User.SemesterMatriculed,
		request.User.GithubUsername,
	)

	userModule, err := s.Storage.GetUserModule()
	if err != nil {
		return err
	}

	if err := userModule.Update(user); err != nil {
		return server.SendError(w, err)
	}

	return server.WriteJSON(w, http.StatusOK, nil)
}

func DeleteUser(s server.Server, w http.ResponseWriter, r *http.Request) error {
	idStr, err := server.GetRequestParam(r, "id")
	if err != nil {
		return server.SendError(w, err)
	}

	id, err := strconv.ParseUint(*idStr, 10, 64)
	if err != nil {
		return server.SendError(w, err)
	}

	userModule, err := s.Storage.GetUserModule()
	if err != nil {
		return err
	}

	err = userModule.Delete(id)
	if err != nil {
		return server.SendError(w, err)
	}

	return server.WriteJSON(w, http.StatusOK, nil)
}
