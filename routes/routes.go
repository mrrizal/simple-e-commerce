package routes

import (
	"e-commerce-api/app/configs"
	"e-commerce-api/app/controllers"
	"e-commerce-api/app/middlewares"
	"e-commerce-api/app/services"
	"e-commerce-api/app/validators"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(app *fiber.App, db *pgxpool.Pool) {
	config := configs.LoadConfig()
	// Initialize services
	customerService := services.NewCustomerService(db)
	productService := services.NewProductService(db)
	orderService := services.NewOrderService(db)

	// initialize validator
	customerValidator := validators.NewCustomerValidator(customerService, config)

	productValidator := validators.ProductValidator{
		ProductService: productService,
	}

	// Initialize controllers
	customerController := &controllers.CustomerController{
		CustomerService: customerService,
	}

	productController := &controllers.ProductController{
		ProductService: productService,
	}

	orderController := &controllers.OrderController{
		OrderService:      orderService,
		ProductService:    productService,
		CustomerValidator: customerValidator,
		ProductValidator:  productValidator,
		Config:            config,
	}

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
