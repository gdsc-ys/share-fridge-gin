package controller

import (
	"net/http"

	"github.com/gdsc-ys/share-fridge-gin/model"
	"github.com/gin-gonic/gin"
)

func ReadUser(c *gin.Context) {
	user := model.User{}

	if err := model.DB.Where("id = ? AND type = ?", c.Param("id"), c.Param("type")).Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func CreateUser(c *gin.Context) {
	user := model.User{}

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func ReadUserFavorite(c *gin.Context) {
	user := model.User{}

	if err := model.DB.Model(&user).Where("id = ? AND type = ?", c.Param("id"), c.Param("type")).Association("Fridges").Find(&user.Fridges); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func UpdateUserFavorite(c *gin.Context) {
	user := model.User{}

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.DB.Model(&user).Where("id = ? AND type = ?", c.Param("id"), c.Param("type")).Association("Fridges").Replace(user.Fridges); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
