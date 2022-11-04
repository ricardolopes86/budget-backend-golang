package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ricardosilva86/budget-backend-golang/models"
)

type CreateEntradaInput struct {
	Valor     float64   `json:"valor" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

func CreateEntrada(c *gin.Context) {
	var input CreateEntradaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Printf("%v\n", input)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entrada := models.Entrada{Valor: input.Valor, CreatedAt: input.CreatedAt}
	models.DB.Create(&entrada)

	c.JSON(http.StatusOK, gin.H{"data": entrada})
}
