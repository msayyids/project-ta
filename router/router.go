package router

import (
	"project-ta/config"
	"project-ta/controller"
	"project-ta/helper"
	"project-ta/middleware"
	"project-ta/repository"
	"project-ta/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

// NewRouter initializes and returns a new router
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	// Initialize database connection
	validate := validator.New()
	db := config.ConnectDb()

	coreAPIClient := config.SetupCoreAPIClient()

	// Initialize repository, service, and controller
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo, db, *validate)
	userController := controller.NewUserController(userService)

	layananRepo := repository.NewLayananRepository()
	layananService := service.NewLayananService(db, layananRepo)
	layananController := controller.NewLayananController(layananService, *validate)

	orderRepo := repository.NewOrderRepository()
	orderService := service.NewOrderService(db, validate, orderRepo, layananRepo)
	orderControllerr := controller.NewOrderController(orderService, *coreAPIClient, *validate)

	paymentRepo := repository.NewPaymentRepository()
	paymentService := service.NewPaymentService(db, validate, paymentRepo, orderRepo, layananRepo)
	paymentController := controller.NewPaymentController(paymentService, orderService, *coreAPIClient)
	_ = paymentController

	pengeluaranRepo := repository.NewPengeluaranRepository()
	pengeluranService := service.NewPengeluaranService(db, pengeluaranRepo)
	pengeluaranController := controller.NewPengeluaranController(pengeluranService)

	keuntunganRepo := repository.NewKeuntunganRepository()
	keuntunganService := service.NewKeuntunganService(db, keuntunganRepo)
	keuntunganController := controller.NewKeuntunganCntroller(keuntunganService)

	// // middleware
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

	router.GET("/api/steam/order", orderControllerr.FindAll)
	router.GET("/api/steam/order/:id", userMiddleware.AuthUser(orderControllerr.FindById))
	router.PUT("/api/steam/order/:id", adminMiddleware.AuthAdmin(orderControllerr.UpdateOrder))

	router.POST("/api/steam/order/cashless", userMiddleware.AuthUser(orderControllerr.CreateOrderCashless))
	router.POST("/api/steam/order/cash", userMiddleware.AuthUser(orderControllerr.CreateOrderCash))
	// router.POST("/api/steam/payment", paymentController.CreatePaymentEmoney)

	router.POST("/api/steam/pengeluaran", pengeluaranController.CreatePengeluaran)
	router.GET("/api/steam/pengeluaran", pengeluaranController.GetPengeluaran)
	router.GET("/api/steam/pengeluaran/:id", pengeluaranController.GetPengeluaranById)
	router.PUT("/api/steam/pengeluaran/:id", pengeluaranController.UpdatePengeluaran)
	router.DELETE("/api/steam/pengeluaran/:id", pengeluaranController.DeletePengeluaran)

	router.GET("/api/steam/keuntungan/hari/:tanggal", adminMiddleware.AuthAdmin(keuntunganController.GetKeuntunganByDateEndpoint))
	router.GET("/api/steam/keuntungan/bulan/:tahun/:bulan", adminMiddleware.AuthAdmin(keuntunganController.GetKeuntunganByMonthEndpoint))
	router.GET("/api/steam/keuntungan/minggu/:tanggal", adminMiddleware.AuthAdmin(keuntunganController.GetKeuntunganByLast7DaysEndpoint))

	router.POST("/webhooks", paymentController.VerifyPayment)

	router.PanicHandler = helper.PanicHandlerWrapper

	return router
}
