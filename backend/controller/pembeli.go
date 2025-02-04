package controller

import (
	"backend/database"
	"backend/models"
	. "backend/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collPembeli = database.Db.Collection("pembeli")

func RegisterPembeli(c *gin.Context){
	ctx := c.Request.Context()
	var pembeli models.Pembeli
	err := c.BindJSON(&pembeli)
	if err != nil {
		c.JSON(http.StatusConflict, Message("gagal bind data"))
		return
	}
	err = validate.Struct(pembeli)
	if err != nil {
		c.JSON(http.StatusBadRequest, Message("format data salah"))
		return
	}
	count, err := collPembeli.CountDocuments(ctx, bson.M{"username": pembeli.Username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, Message("error mencari username"))
		return
	}
	if count > 0 {
		c.JSON(http.StatusFound, Message("Username sudah ada"))
		return
	}
	hashPass, err := HashPassword(pembeli.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Message("gagal hash password"))
		return
	}
	pembeli.Password = hashPass
	_, err = collPembeli.InsertOne(ctx, pembeli)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Message("gagal tambah user"))
		return
	}
	c.IndentedJSON(http.StatusOK, Message("berhasil nambah pembeli", pembeli))
}

func LoginPembeli(c *gin.Context){
	ctx := c.Request.Context()
	var reqPembeli models.Pembeli
	err := c.BindJSON(&reqPembeli)
	if err != nil {
		c.JSON(http.StatusBadRequest, Message("gagal bind request"))
		return
	}
	var pembeli models.Pembeli
	filter := bson.M{"username": reqPembeli.Username}
	err = collPembeli.FindOne(ctx, filter).Decode(&pembeli)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, Message("username tidak ada"))
			return
		}
		c.JSON(http.StatusInternalServerError, Message("gagal mendapatkan user"))
		return
	}
	err = CekPassword(pembeli.Password, reqPembeli.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "password salah")
		return
	}
	session := sessions.Default(c)
	session.Set("username", pembeli.Username)
	session.Set("role", "pembeli")
	session.Save()
	c.IndentedJSON(http.StatusOK, Message("berhasil login pembeli"))
}