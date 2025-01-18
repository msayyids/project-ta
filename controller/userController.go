package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	S service.UserServiceInj
}

type UserControllerInj interface {
	CreateUsers(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	GetUser(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	GetAllUser(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	EditUser(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	DeleteUser(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}

func NewUserController(s service.UserServiceInj) UserControllerInj {
	return UserController{S: s}
}

func (uc UserController) CreateUsers(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	newUser := entity.UserRequest{}

	validate := validator.New()

	helper.RequestBody(r, &newUser)

	err := validate.Struct(newUser)
	if err != nil {

		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "BAD REQUEST",
			Data:    "INVALID INPUT",
		}, http.StatusBadRequest)
		return
	}

	newUser.Password = helper.HashPassword(newUser.Password)

	// Proses pembuatan user
	newUserResponse, err := uc.S.CreateUser(r.Context(), newUser)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "INTERNAL SERVER ERROR",
			Data:    "Failed to create user",
		}, http.StatusInternalServerError)
		return
	}

	// Response sukses
	helper.ResponseBody(w, entity.WebResponse{
		Code:    http.StatusCreated,
		Message: "CREATED",
		Data:    newUserResponse,
	}, http.StatusCreated)
}

func (uc UserController) Login(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	loginRequser := entity.LoginRequest{}
	helper.RequestBody(r, &loginRequser)

	// Find user by email
	loggedInUser, err := uc.S.FindUserByEmail(r.Context(), loginRequser.Email)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    404,
			Message: "NOT FOUND",
			Data:    "Email not registered",
		}, http.StatusNotFound)
		return
	}

	// Validate password
	if !helper.ValidatePassword(loginRequser.Password, loggedInUser.Password) {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    400,
			Message: "BAD REQUEST",
			Data:    "Incorrect password",
		}, http.StatusBadRequest)
		return
	}

	token := helper.CreateJWT(jwt.MapClaims{
		"id":   loggedInUser.ID,
		"role": loggedInUser.Role,
	})

	response := entity.WebResponse{
		Code:    http.StatusOK,
		Message: "success login",
		Data:    token,
	}

	helper.ResponseBody(w, response, http.StatusOK)
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")
	idInt, _ := strconv.Atoi(id)
	user, err := uc.S.FindUserById(r.Context(), idInt)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusNotFound,
			Message: "NOT FOUND",
			Data:    fmt.Sprintf("User with ID %s not found", id),
		}, http.StatusNotFound)
		return
	}

	// Success response
	helper.ResponseBody(w, entity.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    user,
	}, http.StatusOK)
}

func (uc UserController) GetAllUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	users, err := uc.S.FindAllUsers(r.Context())
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
			Data:    "Failed to retrieve users",
		}, http.StatusInternalServerError)
		return
	}

	// Success response
	helper.ResponseBody(w, entity.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    users,
	}, http.StatusOK)
}

func (uc UserController) EditUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")
	idInt, _ := strconv.Atoi(id)

	userReq := entity.UserRequest{}
	helper.RequestBody(r, &userReq)

	// Update user
	_, err := uc.S.EditUser(r.Context(), idInt, userReq)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "BAD REQUEST",
			Data:    "Failed to update user",
		}, http.StatusBadRequest)
		return
	}

	// Success response
	helper.ResponseBody(w, entity.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    "success update data user ",
	}, http.StatusOK)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")
	idInt, _ := strconv.Atoi(id)

	// Delete user
	err := uc.S.DeleteUser(r.Context(), idInt)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:    http.StatusBadRequest,
			Message: "BAD REQUEST",
			Data:    "Failed to delete user",
		}, http.StatusBadRequest)
		return
	}

	helper.ResponseBody(w, entity.WebResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    fmt.Sprintf("User with ID %s has been deleted", id),
	}, http.StatusOK)
}
