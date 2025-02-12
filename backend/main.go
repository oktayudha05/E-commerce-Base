package main

import (
	"backend/controller"
	"backend/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("gagal load .env di main")
	}
}

func main() {
	router := gin.Default()
	router.Use(middleware.SetupSession())
	api := router.Group("/api")

	penjual := api.Group("/auth/penjual")
	{
		penjual.POST("/register", controller.RegisterPenjual)
		penjual.POST("/login", controller.LoginPenjual)
	}
	pembeli := api.Group("/auth/pembeli")
	{
		pembeli.POST("/register", controller.RegisterPembeli)
		pembeli.POST("/login", controller.LoginPembeli)
	}
	router.POST("/barang", middleware.Auth("penjual"), controller.AddBarang)

	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.JSON(200, gin.H{"message": "berhasil hapus session"})
	})

	router.GET("/dashboard", middleware.Auth("penjual"), func(c *gin.Context){
		session := sessions.Default(c)
		username := session.Get("username")
		idpenjual := session.Get("penjual_id")
		c.JSON(200, gin.H{"message": "halo bro " + username.(string) + idpenjual.(string)})
	})
	router.Run("127.0.0.1:3001")
}