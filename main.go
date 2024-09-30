package main

import (
	"github.com/BananaFried525/home-restaurant-api/src/configs"
	"github.com/BananaFried525/home-restaurant-api/src/database"
	"github.com/BananaFried525/home-restaurant-api/src/middlewares"
	"github.com/BananaFried525/home-restaurant-api/src/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.Init()
	database.Init()

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.CustomRecovery(middlewares.Recovery))

	routers.New(r)

	r.Run(":8080")
}
