package router

import (
	"project-ta/config"
	"project-ta/controller"
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
	midtrans := config.SetupMidtrans()

	// Initialize repository, service, and controller
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo, &db)
	userController := controller.NewUserController(userService)

	layananRepo := repository.NewLayananRepository()
	layananService := service.NewLayananService(&db, layananRepo)
	layananController := controller.NewLayananController(layananService)

	orderRepo := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepo, db, layananRepo)
	orderControllerr := controller.NewOrderController(orderService)

	paymentRepo := repository.NewMidtransPaymentRepository()
	paymentService := service.NewPaymentService(*midtrans, db, paymentRepo, orderRepo)
	paymentController := controller.NewPaymentController(paymentService)

	// middleware
	adminMiddleware := middleware.NewAuthAdmin(userService, layananService)
	userMiddleware := middleware.NewAuthUser(userService, layananService)
	// karyawanMiddleware := middleware.NewAuthKaryawan(userService, layananService)

	router.POST("/api/steam/users/login", userController.Login)

	router.POST("/api/steam/users/karyawan", adminMiddleware.AuthAdmin(userController.CreateUsers))
	router.PUT("/api/steam/users/karyawan/:id", adminMiddleware.AuthAdmin(userController.EditUser))
	router.DELETE("/api/steam/users/karyawan/:id", adminMiddleware.AuthAdmin(userController.DeleteUser))
	router.GET("/api/steam/users/karyawan/:id", adminMiddleware.AuthAdmin(userController.GetUser))
	router.GET("/api/steam/users/karyawan/", adminMiddleware.AuthAdmin(userController.GetAllUser))

	router.POST("/api/steam/layanan", adminMiddleware.AuthAdmin(layananController.CreateLayanan))
	router.GET("/api/steam/layanan", userMiddleware.AuthUser(layananController.FindAllLayanan))
	router.DELETE("/api/steam/layanan/:id", adminMiddleware.AuthAdmin(layananController.DeleteLayananById))
	router.GET("/api/steam/layanan/:id", userMiddleware.AuthUser(layananController.FindLayananById))
	router.PUT("/api/steam/layanan/:id", adminMiddleware.AuthAdmin(layananController.EditLayananById))

	router.POST("/api/steam/order", userMiddleware.AuthUser(orderControllerr.CreateOrder))
	router.GET("/api/steam/order", userMiddleware.AuthUser(orderControllerr.GetAllOrder))
	router.GET("/api/steam/order/:id", userMiddleware.AuthUser(orderControllerr.GetOrderById))
	router.PUT("/api/steam/order/:id", adminMiddleware.AuthAdmin(orderControllerr.EditOrder))
	router.DELETE("/api/steam/order/:id", adminMiddleware.AuthAdmin(orderControllerr.DeleteORder))

	router.POST("/api/steam/payment", paymentController.CreatePayment)

	// router.PanicHandler = helper.ErrorHandler

	return router
}
