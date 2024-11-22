package router

import (
	"project-ta/config"
	"project-ta/controller"
	"project-ta/helper"
	"project-ta/repository"
	"project-ta/service"

	"github.com/julienschmidt/httprouter"
)

// NewRouter initializes and returns a new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	// Initialize database connection
	db := config.ConnecDb()

	// Initialize repository, service, and controller
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo, &db)
	userController := controller.NewUserController(userService)

	// Public Endpoints
	router.POST("/cucimotorberkah/api/users", userController.CreateUsers)
	router.POST("/cucimotorberkah/api/users/login", userController.Login)

	// Error handling
	router.PanicHandler = helper.ErrorHandler

	return router
}
