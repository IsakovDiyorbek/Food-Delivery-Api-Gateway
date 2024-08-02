package api

import (
	"github.com/Food_Delivery/Food-Delivery-Api-Gateway/api/handler"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewGin(users, delivery *grpc.ClientConn) *gin.Engine{
	h := handler.NewHandler(users, delivery)
	r := gin.Default()
	r.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/product", h.CreateProduct)
	r.GET("/product/:id", h.GetProduct)
	r.PUT("/product/:id", h.UpdateProduct)
	r.DELETE("/product/:id", h.DeleteProduct)
	r.GET("/products", h.GetAllProducts)

	r.POST("/task", h.CreateTask)
	r.GET("/task/:id", h.GetTask)
	r.PUT("/task/:id", h.UpdateTask)
	r.DELETE("/task/:id", h.DeleteTask)
	r.GET("/tasks", h.GetAllTasks)

	r.POST("/courier_location", h.CreateCourierLocation)
	r.GET("/courier_location/:id", h.GetCourierLocation)
	r.PUT("/courier_location/:id", h.UpdateCourierLocation)
	r.DELETE("/courier_location/:id", h.DeleteCourierLocation)
	r.GET("/courier_locations", h.GetAllCourierLocations)


	r.POST("/order", h.CreateOrder)
	r.GET("/order/:id", h.GetOrder)
	r.PUT("/order/:id", h.UpdateOrder)
	r.DELETE("/order/:id", h.DeleteOrder)
	r.GET("/orders", h.GetAllOrders)

	r.POST("/order_item", h.CreateOrderItem)
	r.GET("/order_item/:id", h.GetOrderItem)
	r.PUT("/order_item/:id", h.UpdateOrderItem)
	r.DELETE("/order_item/:id", h.DeleteOrderItem)
	r.GET("/order_items", h.GetAllOrderItems)



	return r


}
