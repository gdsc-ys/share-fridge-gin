package controller

import (
	"net/http"

	"github.com/gdsc-ys/share-fridge-gin/model"
	"github.com/gin-gonic/gin"
)

func ReadFridge(c *gin.Context) {
	fridge := model.Fridge{}

	if err := model.DB.First(&fridge, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": fridge,
	})
}

func CreateFridge(c *gin.Context) {
	fridge := model.Fridge{}
	if err := c.ShouldBind(&fridge); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	model.DB.Create(&fridge)
	c.JSON(http.StatusOK, gin.H{
		"data": fridge,
	})
}
