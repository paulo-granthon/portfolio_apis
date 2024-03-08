package endpoints

import (
	"models"
	"net/http"
	"schemas"
	"server"
	"strconv"

	"github.com/ztrue/tracerr"
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
		{
			Path: "/users/{id}/projects",
			Methods: []server.Method{
				server.NewMethod("GET", GetUserProjects),
			},
		},
	}
}

func GetUsers(s server.Server, w http.ResponseWriter, r *http.Request) error {
	userModule, err := s.Storage.GetUserModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	users, err := userModule.Get()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error getting users",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		users,
	)
}

func GetUser(s server.Server, w http.ResponseWriter, r *http.Request) error {
	idStr, err := server.GetRequestParam(r, "id")
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter id not found",
		)
	}

	id, err := strconv.ParseUint(*idStr, 10, 64)
	if err != nil {
		return server.SendError(
			w, tracerr.Wrap(err), http.StatusBadRequest,
			"Parameter id is not a valid number",
		)
	}

	userModule, err := s.Storage.GetUserModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	user, err := userModule.GetById(id)
	if err != nil {
		return server.SendError(
			w, err, http.StatusNotFound,
			"user not found",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		user,
	)
}

func RegisterUser(s server.Server, w http.ResponseWriter, r *http.Request) error {
	var request schemas.RegisterUserRequest
	if err := server.ReadJSON(r, &request.Payload); err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"error reading request",
		)
	}

	user := models.NewRegisterUser(
		request.Payload.Name,
		request.Payload.Password,
	)

	userModule, err := s.Storage.GetUserModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	id, err := userModule.Register(user)
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error registering user",
		)
	}

	return server.WriteJSON(
		w, http.StatusCreated,
		schemas.RegisterUserResponse{Id: *id},
	)
}

func CreateUser(s server.Server, w http.ResponseWriter, r *http.Request) error {
	var request schemas.CreateUserRequest
	if err := server.ReadJSON(r, &request.User); err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"error parsing request",
		)
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
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	id, err := userModule.Create(user)
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error creating user",
		)
	}

	return server.WriteJSON(
		w, http.StatusCreated,
		schemas.CreateUserResponse{Id: *id},
	)
}

func UpdateUser(s server.Server, w http.ResponseWriter, r *http.Request) error {
	idStr, err := server.GetRequestParam(r, "id")
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter id not found",
		)
	}

	id, err := strconv.ParseUint(*idStr, 10, 64)
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter id is not a valid number",
		)
	}

	var request schemas.UpdateUserRequest
	if err := server.ReadJSON(r, &request); err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"error parsing request",
		)
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
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	if err := userModule.Update(user); err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error updating user",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		"user updated",
	)
}

func DeleteUser(s server.Server, w http.ResponseWriter, r *http.Request) error {
	idStr, err := server.GetRequestParam(r, "id")
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter id not found",
		)
	}

	id, err := strconv.ParseUint(*idStr, 10, 64)
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter id is not a valid number",
		)
	}

	userModule, err := s.Storage.GetUserModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	err = userModule.Delete(id)
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error deleting user",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		"user deleted",
	)
}

func GetUserProjects(s server.Server, w http.ResponseWriter, r *http.Request) error {
	idStr, err := server.GetRequestParam(r, "id")
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter id not found",
		)
	}

	id, err := strconv.ParseUint(*idStr, 10, 64)
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter id is not a valid number",
		)
	}

	projectModule, err := s.Storage.GetProjectModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	projects, err := projectModule.GetByUserId(id)
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error getting projects",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		projects,
	)
}
