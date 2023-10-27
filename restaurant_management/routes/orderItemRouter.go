package routes

import (
	"restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderitems", controllers.GetOrderItems())
	incomingRoutes.GET("/ordersitems/:orderitem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/ordersitems-order/:order_id", controllers.GetOrderItemByOrder())
	incomingRoutes.POST("/orderitems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderitems/:orderitem_id", controllers.UpdateOrderItem())
}
