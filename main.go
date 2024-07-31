package main

import (
	"log"

	"github.com/Food_Delivery/Food-Delivery-Api-Gateway/api"
	_ "github.com/Food_Delivery/Food-Delivery-Api-Gateway/docs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	delivery, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer delivery.Close()

	user, err := grpc.NewClient(":8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer user.Close()

	r := api.NewGin(user, delivery)
	log.Println("starting server")
	r.Run(":8088")
	
}
