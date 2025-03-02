package controller

import (
	"backend/database"
	"backend/models"
	"backend/utils"
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
		c.JSON(http.StatusConflict, utils.Message("gagal bind data"))
		return
	}
	err = validate.Struct(pembeli)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Message("format data salah"))
		return
	}
	count, err := collPembeli.CountDocuments(ctx, bson.M{"username": pembeli.Username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Message("error mencari username"))
		return
	}
	if count > 0 {
		c.JSON(http.StatusFound, utils.Message("Username sudah ada"))
		return
	}
	hashPass, err := utils.HashPassword(pembeli.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Message("gagal hash password"))
		return
	}
	pembeli.Password = hashPass
	_, err = collPembeli.InsertOne(ctx, pembeli)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Message("gagal tambah user"))
		return
	}
	c.IndentedJSON(http.StatusOK, utils.Message("berhasil nambah pembeli", pembeli))
}

func LoginPembeli(c *gin.Context){
	ctx := c.Request.Context()
	var reqPembeli struct {
		Username string `json:"username"`
		Password string `json:"Password"`
	}
	err := c.BindJSON(&reqPembeli)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Message("gagal bind request"))
		return
	}
	err = validate.Struct(reqPembeli)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Message("data tidak sesuai format"))
		return
	}
	var pembeli models.Pembeli
	filter := bson.M{"username": reqPembeli.Username}
	err = collPembeli.FindOne(ctx, filter).Decode(&pembeli)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, utils.Message("username tidak ada"))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.Message("gagal mendapatkan user"))
		return
	}
	err = utils.CekPassword(pembeli.Password, reqPembeli.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "password salah")
		return
	}
	session := sessions.Default(c)
	session.Set("username", pembeli.Username)
	session.Set("role", "pembeli")
	session.Save()
	c.IndentedJSON(http.StatusOK, utils.Message("berhasil login pembeli"))
}