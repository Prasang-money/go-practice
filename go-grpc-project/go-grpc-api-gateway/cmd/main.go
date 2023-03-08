package main

import (
	"log"
	"os"

	"github.com/Prasang-money/go-grpc-api-gateway/pkg/auth"
	"github.com/Prasang-money/go-grpc-api-gateway/pkg/config"
	"github.com/Prasang-money/go-grpc-api-gateway/pkg/order"
	"github.com/Prasang-money/go-grpc-api-gateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	if err = r.Run(c.Port); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
