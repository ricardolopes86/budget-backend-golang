package main

import (
	"github.com/ricardosilva86/budget-backend-golang/controllers"
	"github.com/ricardosilva86/budget-backend-golang/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/entradas", controllers.CreateEntrada)
	router.GET("/entradas", controllers.FindEntradas)
	router.GET("/entradas/:id", controllers.FindEntrada)
	router.DELETE("/entradas/:id", controllers.DeleteEntrada)
	router.PATCH("/entradas/:id", controllers.UpdateEntrada)
	router.Run("localhost:9000")
}
