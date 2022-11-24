package main

import (
	"github.com/ricardosilva86/budget-backend-golang/controllers"
	"github.com/ricardosilva86/budget-backend-golang/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	models.ConnectDatabase()

	router.POST("/api/v1/entradas", controllers.CreateEntrada)
	router.GET("/api/v1/entradas", controllers.FindEntradas)
	router.GET("/api/v1/entradas/:id", controllers.FindEntrada)
	router.DELETE("/api/v1/entradas/:id", controllers.DeleteEntrada)
	router.PATCH("/api/v1/entradas/:id", controllers.UpdateEntrada)

	router.POST("/api/v1/fixos", controllers.CreateFixos)
	router.GET("/api/v1/fixos", controllers.FindFixos)
	router.GET("/api/v1/fixos/:id", controllers.FindFixo)
	router.DELETE("/api/v1/fixos/:id", controllers.DeleteFixo)
	router.PATCH("/api/v1/fixos/:id", controllers.UpdateFixo)

	router.Run("localhost:9000")
}
