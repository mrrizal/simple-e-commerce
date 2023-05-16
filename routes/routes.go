package routes

import (
	// "project/app/controllers"
	// "project/app/services"

	"tabungan-api/app/controllers"
	"tabungan-api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(app *fiber.App, db *pgxpool.Pool) {
	// Initialize services
	customerService := services.NewCustomerService(db)
	productService := services.NewProductService(db)

	// Initialize controllers
	customerController := &controllers.CustomerController{
		CustomerService: customerService,
	}

	productController := &controllers.ProductController{
		ProductService: productService,
	}

	// Set up routes
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/customer/sign-up/", customerController.SingUp)
	v1.Post("/customer/sign-in/", customerController.SignIn)
	v1.Get("/products/:id/", productController.Get)
	v1.Get("/products/", productController.GetAll)
	v1.Post("/products/multiple/", productController.GetMultiple)
}
