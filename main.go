package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gdsc-ys/share-fridge-gin/controller"
	"github.com/gdsc-ys/share-fridge-gin/model"
)

func main() {
	router := gin.Default()

	model.ConnectDatabase()

	user := router.Group("/user")
	{
		user.GET("/:type/:id", controller.ReadUser)
		user.POST("", controller.CreateUser)
		user.PUT("/:type/:id/favorite", controller.UpdateUserFavorite)
	}

	food := router.Group("/food")
	{
		food.GET("/:id", controller.ReadFood)
		food.POST("", controller.CreateFood)
	}

	fridge := router.Group("/fridge")
	{
		fridge.GET("/:id", controller.ReadFridge)
		fridge.POST("", controller.CreateFridge)
	}

	router.Run(":8080")
}
