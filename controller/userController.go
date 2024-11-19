package controller

import (
	"fmt"
	"net/http"
	"project-ta/entity"
	"project-ta/helper"
	"project-ta/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	S service.UserServiceInj
}

type UserControllerInj interface {
	CreateUsers(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, param httprouter.Params)
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
	loggedInUser, err := uc.S.FindUSerByemail(r.Context(), loginRequser.Email)
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
