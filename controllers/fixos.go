package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardosilva86/budget-backend-golang/models"
)

type CreateFixosInput struct {
	Contas      json.Number `json:"contas"`
	Assinaturas json.Number `json:"assinaturas"`
	Seguros     json.Number `json:"seguros"`
	Mesadas     json.Number `json:"mesadas"`
	Impostos    json.Number `json:"impostos"`
	Outros      json.Number `json:"outros"`
	Mes         json.Number `json:"mes" binding:"required"`
	Ano         json.Number `json:"ano" binding:"required"`
}

type UpdateFixosInput struct {
	Contas      json.Number `json:"contas"`
	Assinaturas json.Number `json:"assinaturas"`
	Seguros     json.Number `json:"seguros"`
	Mesadas     json.Number `json:"mesadas"`
	Impostos    json.Number `json:"impostos"`
	Outros      json.Number `json:"outros"`
	Mes         json.Number `json:"mes" binding:"required"`
	Ano         json.Number `json:"ano" binding:"required"`
}

func CreateFixos(c *gin.Context) {
	var input CreateFixosInput
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Printf("%v\n", input)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("creating Fixo")
	Fixo := models.Fixos{Contas: input.Contas, Assinaturas: input.Assinaturas, Seguros: input.Seguros, Mesadas: input.Mesadas, Impostos: input.Impostos, Outros: input.Outros, Mes: input.Mes, Ano: input.Ano}
	models.DB.Create(&Fixo)

	c.JSON(http.StatusOK, gin.H{"data": Fixo})
}

func FindFixos(c *gin.Context) {
	var Fixos []models.Fixos
	models.DB.Find(&Fixos)

	c.JSON(http.StatusOK, gin.H{"data": Fixos})
}

func FindFixo(c *gin.Context) {
	var Fixo models.Fixos

	if err := models.DB.Where("id = ?", c.Param("id")).First(&Fixo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Fixo})
}

func DeleteFixo(c *gin.Context) {
	var Fixo models.Fixos
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Fixo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(&Fixo)
	c.JSON(http.StatusOK, gin.H{"data": "data deleted"})
}

func UpdateFixo(c *gin.Context) {
	var Fixo models.Fixos
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Fixo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	var input UpdateFixosInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedFixo := models.Fixos{Contas: input.Contas, Assinaturas: input.Assinaturas, Seguros: input.Seguros, Mesadas: input.Mesadas, Impostos: input.Impostos, Outros: input.Outros, Mes: input.Mes, Ano: input.Ano}
	models.DB.Model(&Fixo).Updates(&updatedFixo)
	c.JSON(http.StatusOK, gin.H{"data": Fixo})
}
