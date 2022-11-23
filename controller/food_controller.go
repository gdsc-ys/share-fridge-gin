package controller

import (
	"net/http"

	"github.com/gdsc-ys/share-fridge-gin/model"
	"github.com/gin-gonic/gin"
)

func ReadFood(c *gin.Context) {
	food := model.Food{}

	if err := model.DB.First(&food, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": food,
	})
}

func CreateFood(c *gin.Context) {
	food := model.Food{}
	if err := c.ShouldBind(&food); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	model.DB.Create(&food)
	c.JSON(http.StatusOK, gin.H{
		"data": food,
	})
}
