package main

import (
	"app/config"
	"app/controllers"
	"app/datasources"
	"app/deliveries"
	"app/managers"
)

func main() {
	config.LoadConfig("config/config.yaml")

	restClient := datasources.NewRestClient()
	redisClient := datasources.NewRedisCache()
	adsManager := managers.NewAdvertiserManager(restClient, redisClient)
	adsController := controllers.NewAdsController(adsManager)

	deliveries.StartHTTPServer(adsController)
}
