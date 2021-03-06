package main

import (
	"fileserver/api"
	"fileserver/auth"
	"fileserver/config"
	"fileserver/db"
	"fileserver/fs"
	"fileserver/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.New()

	db.New(cfg.DB)
	fs.New(cfg.Root)
	user.New()
	auth.New(cfg.Casbin.Model)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))
	router.Use(gin.Logger())

	api.NewRouter(router)

	router.Run(cfg.Port)
}
