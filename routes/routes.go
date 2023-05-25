package routes

import (
	"e-commerce-api/app/controllers"
	"e-commerce-api/app/database"
	"e-commerce-api/app/middlewares"
	"e-commerce-api/app/services"
	"e-commerce-api/app/validators"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db database.DB) {
	// Initialize services
	customerService := services.NewCustomerService(db)
	productService := services.NewProductService(db)
	orderService := services.NewOrderService(db)

	// initialize validator
	customerValidator := validators.NewCustomerValidator(customerService)
	productValidator := validators.NewProductValidator(productService)

	// Initialize controllers
	customerController := controllers.NewCustomerController(customerService)
	productController := controllers.NewProductController(productService)
	orderController := controllers.NewOrderController(orderService, productService,
		customerValidator, productValidator)

	// Set up routes
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/customer/sign-up/", customerController.SingUp)
	v1.Post("/customer/sign-in/", customerController.SignIn)
	v1.Get("/products/:id/", productController.Get)
	v1.Get("/products/", productController.GetAll)
	v1.Post("/products/multiple/", productController.GetMultiple)
	v1.Post("/order/", middlewares.JWTMiddleware, orderController.CreateOrder)
	v1.Get("/order/", middlewares.JWTMiddleware, orderController.Get)
}
