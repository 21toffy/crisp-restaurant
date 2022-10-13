package routes

import (
	controller "github.com/21toffy/crisp-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/foods", controller.GetFoods())
	incommingRoutes.GET("/foods/:food_id", controller.GetFood())
	incommingRoutes.POST("/food", controller.CreateFood())
	incommingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())
}
