package main

import (
	"fmt"
	"log"
	"tabungan-api/app/configs"
	"tabungan-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AppInstance struct {
	app *fiber.App
	db  *pgxpool.Pool
}

func NewAppInstance(config configs.Config) AppInstance {
	db, err := configs.NewDB(config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	return AppInstance{app: app, db: db}
}

func main() {
	config := configs.LoadConfig()
	app := NewAppInstance(config)

	routes.SetupRoutes(app.app, app.db)
	app.app.Listen(fmt.Sprintf(":%s", config.Port))
}
