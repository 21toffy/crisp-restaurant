package routes

import (
	controller "github.com/21toffy/crisp-restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/orderItems", controller.GetOrderItems())
	incommingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incommingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incommingRoutes.GET("/orderItems-order/order_id", controller.GetOrderItemsByOrder())
	incommingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())

}
