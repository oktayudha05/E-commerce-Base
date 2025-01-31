package main

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	penjual := router.Group("/auth/penjual")
	{
		penjual.POST("/register", controller.RegisterPenjual)
		penjual.POST("/login", controller.LoginPenjual)
	}
	pembeli := router.Group("/auth/pembeli")
	{
		pembeli.POST("/register", controller.RegisterPembeli)
		pembeli.POST("/login", controller.LoginPembeli)
	}
	router.Run(":3001")
}