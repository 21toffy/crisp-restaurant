package routes

import (
	controller "github.com/21toffy/crisp-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/orders", controller.GetOrders())
	incommingRoutes.GET("/orders/:order_id", controller.GetOrder())
	incommingRoutes.POST("/orders", controller.CreateOrder())
	incommingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())

}
