package controller

import (
	"fmt"
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
	helper.RequestBody(r, &newUser)

	// Hash the password
	hashedPassword := helper.HashPassword(newUser.Password)
	newUser.Password = hashedPassword

	// Create the new user
	newUserResponse, err := uc.S.CreateUser(r.Context(), newUser)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   "Failed to create user",
		})
		return
	}

	// Success response
	response := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   fmt.Sprintf("Welcome %s %s", newUserResponse.Nama_depan, newUserResponse.Nama_belakang),
	}

	helper.ResponseBody(w, response)
}

func (uc UserController) Login(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	loginRequser := entity.UserLoginRequest{}
	helper.RequestBody(r, &loginRequser)

	// Find user by email
	loggedInUser, err := uc.S.FindUserByEmail(r.Context(), loginRequser.Email)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   404,
			Status: "NOT FOUND",
			Data:   "Email not registered",
		})
		return
	}

	// Validate password
	if !helper.ValidatePassword(loginRequser.Password, loggedInUser.Password) {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Incorrect password",
		})
		return
	}

	// Create JWT token
	token := helper.CreateJWT(jwt.MapClaims{
		"id":   loggedInUser.Id,
		"role": loggedInUser.Role,
	})

	response := entity.WebResponse{
		Code:   200,
		Status: "success login",
		Data:   token,
	}

	helper.ResponseBody(w, response)
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")
	idInt, _ := strconv.Atoi(id)
	user, err := uc.S.FindUserById(r.Context(), idInt)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   404,
			Status: "NOT FOUND",
			Data:   fmt.Sprintf("User with ID %s not found", id),
		})
		return
	}

	// Success response
	helper.ResponseBody(w, entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   user,
	})
}

func (uc UserController) GetAllUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	users, err := uc.S.FindAllUsers(r.Context())
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   "Failed to retrieve users",
		})
		return
	}

	// Success response
	helper.ResponseBody(w, entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   users,
	})
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
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Failed to update user",
		})
		return
	}

	// Success response
	helper.ResponseBody(w, entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "success update data user ",
	})
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")
	idInt, _ := strconv.Atoi(id)

	// Delete user
	err := uc.S.DeleteUser(r.Context(), idInt)
	if err != nil {
		helper.ResponseBody(w, entity.WebResponse{
			Code:   400,
			Status: "BAD REQUEST",
			Data:   "Failed to delete user",
		})
		return
	}

	// Success response
	helper.ResponseBody(w, entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   fmt.Sprintf("User with ID %s has been deleted", id),
	})
}
