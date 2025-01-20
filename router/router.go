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

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	validate := validator.New()
	db := config.ConnectDb()
	cld := config.InitializeCloudinary()

	coreAPIClient := config.SetupCoreAPIClient()
	snapAPIClient := config.SetupSnapAPIClient()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo, db, *validate)
	userController := controller.NewUserController(userService)

	layananRepo := repository.NewLayananRepository()
	layananService := service.NewLayananService(db, layananRepo)
	layananController := controller.NewLayananController(layananService, *validate)

	orderRepo := repository.NewOrderRepository()
	orderService := service.NewOrderService(db, validate, orderRepo, layananRepo)
	orderController := controller.NewOrderController(orderService, *snapAPIClient, *validate)

	paymentRepo := repository.NewPaymentRepository()
	paymentService := service.NewPaymentService(db, validate, paymentRepo, orderRepo, layananRepo)
	paymentController := controller.NewPaymentController(paymentService, orderService, *coreAPIClient)

	pengeluaranRepo := repository.NewPengeluaranRepository()
	pengeluranService := service.NewPengeluaranService(db, pengeluaranRepo)
	pengeluaranController := controller.NewPengeluaranController(pengeluranService, cld)

	keuntunganRepo := repository.NewKeuntunganRepository()
	keuntunganService := service.NewKeuntunganService(db, keuntunganRepo)
	keuntunganController := controller.NewKeuntunganCntroller(keuntunganService)

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

	router.GET("/api/steam/order", orderController.FindAll)
	router.GET("/api/steam/order/id/:id", userMiddleware.AuthUser(orderController.FindById))
	router.PUT("/api/steam/order/id/:id", adminMiddleware.AuthAdmin(orderController.UpdateOrder))
	router.GET("/api/steam/order/status/:status", userMiddleware.AuthUser(orderController.FindByStatus))

	router.POST("/api/steam/order/cashless", userMiddleware.AuthUser(orderController.CreateOrderCashless))
	router.POST("/api/steam/order/cash", userMiddleware.AuthUser(orderController.CreateOrderCash))
	// router.POST("/api/steam/payment", paymentController.CreatePaymentEmoney)

	router.GET("/api/steam/pengeluaran/hari/:tanggal", userMiddleware.AuthUser(pengeluaranController.FindPengeluaranByDate))
	router.POST("/api/steam/pengeluaran", userMiddleware.AuthUser(pengeluaranController.CreatePengeluaran))
	router.GET("/api/steam/pengeluaran", userMiddleware.AuthUser(pengeluaranController.GetPengeluaran))
	router.GET("/api/steam/pengeluaran/id/:id", userMiddleware.AuthUser(pengeluaranController.GetPengeluaranById))
	router.PUT("/api/steam/pengeluaran/:id", adminMiddleware.AuthAdmin(pengeluaranController.UpdatePengeluaran))
	router.DELETE("/api/steam/pengeluaran/:id", adminMiddleware.AuthAdmin(pengeluaranController.DeletePengeluaran))

	router.GET("/api/steam/keuntungan/hari/:tanggal", adminMiddleware.AuthAdmin(keuntunganController.GetKeuntunganByDateEndpoint))
	router.GET("/api/steam/keuntungan/bulan/:tahun/:bulan", adminMiddleware.AuthAdmin(keuntunganController.GetKeuntunganByMonthEndpoint))
	router.GET("/api/steam/keuntungan/minggu/:tanggal", adminMiddleware.AuthAdmin(keuntunganController.GetKeuntunganByLast7DaysEndpoint))

	router.POST("/webhooks", paymentController.VerifyPayment)

	router.PanicHandler = helper.PanicHandlerWrapper

	return router
}
