package routes

import (
	"e-commerce-api/app/controllers"
	"e-commerce-api/app/database"
	"e-commerce-api/app/middlewares"
	"e-commerce-api/app/repositories"
	"e-commerce-api/app/restapi"
	"e-commerce-api/app/services"
	"e-commerce-api/app/validators"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db database.DB) {
	// initialize repository
	customerRepository := repositories.NewCustomerRepository(db)
	productRepository := repositories.NewProductRepository(db)
	orderRepository := repositories.NewOrderRepository(db)

	// Initialize services
	customerService := services.NewCustomerService(customerRepository)
	productService := services.NewProductService(productRepository)
	orderService := services.NewOrderService(orderRepository)

	// initialize validator
	customerValidator := validators.NewCustomerValidator(customerService)
	productValidator := validators.NewProductValidator(productService)

	// Initialize controllers
	customerController := controllers.NewCustomerController(customerService)
	productController := controllers.NewProductController(productService)
	orderController := controllers.NewOrderController(orderService, productService,
		customerValidator, productValidator)

	// initialize handler for rest api
	customerHandler := restapi.NewCustomerHandler(customerController)
	orderHandler := restapi.NewOrderHandler(orderController)
	productHandler := restapi.NewProductHandler(productController)

	// Set up routes
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/customer/sign-up/", customerHandler.SignUp)
	v1.Post("/customer/sign-in/", customerHandler.SignIn)
	v1.Get("/products/:id/", productHandler.Get)
	v1.Get("/products/", productHandler.GetAll)
	v1.Post("/products/multiple/", productHandler.GetMultiple)
	v1.Post("/order/", middlewares.JWTMiddleware, orderHandler.CreateOrder)
	v1.Get("/order/", middlewares.JWTMiddleware, orderHandler.Get)
}
