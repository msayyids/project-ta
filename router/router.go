package router

import (
	"project-ta/config"
	"project-ta/controller"
	"project-ta/helper"
	"project-ta/middleware"
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

	layananRepo := repository.NewLayananRepository()
	layananService := service.NewLayananService(&db, layananRepo)
	layananController := controller.NewLayananController(layananService)

	// middleware
	adminMiddleware := middleware.NewAuthAdmin(userService, layananService)
	userMiddleware := middleware.NewAuthUser(userService, layananService)
	karyawanMiddleware := middleware.NewAuthKaryawan(userService, layananService)

	// Public Endpoints

	router.POST("/api/steam/users/login", userController.Login)

	router.POST("/api/steam/users", adminMiddleware.AuthAdmin(userController.CreateUsers))
	router.POST("api/steam/layanan", adminMiddleware.AuthAdmin(layananController.CreateLayanan))
	router.GET("api/steam/layanan", userMiddleware.AuthUser(layananController.FindAllLayanan))
	router.DELETE("api/steam/layanan/:id", adminMiddleware.AuthAdmin(layananController.DeleteLayananById))
	router.GET("api/steam/layanan/:id", layananController.FindLayananById)
	router.PUT("api/steam/layanan/:id", adminMiddleware.AuthAdmin(layananController.EditLayananById))

	// Error handling
	router.PanicHandler = helper.ErrorHandler

	return router
}
