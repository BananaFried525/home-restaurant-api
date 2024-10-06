package main

import (
	"fmt"

	"github.com/BananaFried525/home-restaurant-api/boots"
	"github.com/BananaFried525/home-restaurant-api/internal/core/middlewares"
	"github.com/BananaFried525/home-restaurant-api/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	b := boots.New()

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.CustomRecovery(middlewares.Recovery))

	routers.New(r, b.Database)

	r.Run(fmt.Sprintf(":%s", boots.Config.Port))
}
