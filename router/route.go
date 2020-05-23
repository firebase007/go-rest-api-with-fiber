package router


import (
	"github.com/firebase007/go-rest-api-with-fiber/handler"

	"github.com/firebase007/go-rest-api-with-fiber/middleware"

	"github.com/gofiber/fiber"

	"github.com/gofiber/logger"
)

// SetupRoutes func
func SetupRoutes (app *fiber.App) {
	
	// Middleware
	api := app.Group("/api", logger.New(), middleware.AuthReq())  
	
	// routes
	api.Get("/", handler.GetAllProducts)
	api.Get("/:id", handler.GetSingleProduct)
	api.Post("/", handler.CreateProduct)
	api.Delete("/:id", handler.DeleteProduct)

}