package handler

import (
	"github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto"
	"github.com/Food_Delivery/Food-Delivery-Api-Gateway/genproto/user"
	"google.golang.org/grpc"
)


type Handler struct{
	Product genproto.ProductServiceClient
	User user.UserServiceClient
	Order genproto.OrderServiceClient
	OrderItems genproto.OrderItemServiceClient
	Cart genproto.CartServiceClient
}

func NewHandler(users, delivery *grpc.ClientConn) *Handler{
	return &Handler{
		Product: genproto.NewProductServiceClient(delivery),
		User: user.NewUserServiceClient(users),
		Order: genproto.NewOrderServiceClient(delivery),
		OrderItems: genproto.NewOrderItemServiceClient(delivery),
		Cart: genproto.NewCartServiceClient(delivery),

	}
}

