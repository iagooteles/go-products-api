package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if (err != nil) {
		panic(err)
	}

	// repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	// controller
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "working",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
