package router

import (
	"project-ta/config"
	"project-ta/controller"
	"project-ta/helper"
	"project-ta/repository"
	"project-ta/service"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	db := config.ConnecDb()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo, &db)
	userController := controller.NewUserController(userService)

	router.POST("/cucimotorberkah/api/users", userController.CreateUsers)
	router.POST("/cucimotorberkah/api/users/login", userController.Login)

	router.PanicHandler = helper.ErrorHandler

	return router

}
