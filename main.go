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

	router.Run("localhost:9000")
}
