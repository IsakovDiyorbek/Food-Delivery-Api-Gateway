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



	return r


}
