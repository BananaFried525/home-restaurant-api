package main

import (
	"fmt"

	"github.com/BananaFried525/home-restaurant-api/boots"
	"github.com/BananaFried525/home-restaurant-api/internal/adapters/handler/http/routers"
	"github.com/BananaFried525/home-restaurant-api/internal/core/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	b := boots.New()

	r := gin.New()

	log := middlewares.NewLoggingMiddleware()
	r.Use(log.SystemLog())

	recover := middlewares.NewRecoverMiddleware()
	r.Use(gin.CustomRecovery(recover.Recovery))

	routers.New(r, b.Database)

	r.Run(fmt.Sprintf(":%s", boots.Config.Port))
}
