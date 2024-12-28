package main

import (
	db "go-sqlite-crud-project/config"
	"go-sqlite-crud-project/controller"
	"go-sqlite-crud-project/repository"
	"go-sqlite-crud-project/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db.InitializeDatabase()

	// Create repository, service, and controller for Products
	productRepo := repository.NewProductRepository(db.GetDB())
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	// Initialize Gin router
	r := gin.Default()

	// Routes for Products
	r.POST("/products", productController.CreateProduct)
	r.GET("/products/:id", productController.GetProduct)
	r.GET("/products", productController.GetAllProducts)
	r.PUT("/products/:id", productController.UpdateProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)

	// Start server
	r.Run(":8080")
}
