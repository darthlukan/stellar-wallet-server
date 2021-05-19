package main

import (
	api "github.com/darthlukan/stellar-wallet-server/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", api.Ping)
	router.GET("/healthz", api.Health)

	router.GET("/account/:address", api.GetAccount)
	router.POST("/account", api.CreateAccount)

	router.Run()
}
