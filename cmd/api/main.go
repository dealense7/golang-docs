package main

import (
	"github.com/dealense7/market-price-go/config"
	"github.com/dealense7/market-price-go/routes"
	"github.com/dealense7/market-price-go/utils"
)

func main() {
	config.InitConfig()
	utils.ConnectDatabase()
	routes.RegisterRoutes()
}
