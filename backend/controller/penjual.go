package controller

import (
	"backend/database"
	"backend/models"
	. "backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collPenjual = database.Db.Collection("penjual")
var validate = validator.New()

func RegisterPenjual(c *gin.Context){
	ctx := c.Request.Context()
	var penjual models.Penjual
	err := c.BindJSON(&penjual)
	if err != nil {
		c.JSON(http.StatusConflict, Message("gagal bind data"))
		return
	}
	err = validate.Struct(penjual)
	if err != nil {
		c.JSON(http.StatusBadRequest, Message("format data salah"))
		return
	}
	count, err := collPenjual.CountDocuments(ctx, bson.M{"username": penjual.Username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, Message("error mencari username"))
		return
	}
	if count > 0 {
		c.JSON(http.StatusFound, Message("Username sudah ada"))
		return
	}
	hashPass, err := HashPassword(penjual.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Message("gagal hash password"))
		return
	}
	penjual.Password = hashPass
	_, err = collPenjual.InsertOne(ctx, penjual)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Message("gagal tambah user"))
		return
	}
	c.IndentedJSON(http.StatusOK, Message("berhasil nambah penjual", penjual))
}

func LoginPenjual(c *gin.Context){
	ctx := c.Request.Context()
	var reqPenjual models.Penjual
	err := c.BindJSON(&reqPenjual)
	if err != nil {
		c.JSON(http.StatusBadRequest, Message("gagal bind request"))
		return
	}
	var penjual models.Penjual
	filter := bson.M{"username": reqPenjual.Username, "password": reqPenjual.Password}
	err = collPenjual.FindOne(ctx, filter).Decode(&penjual)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, Message("username atau password salah"))
			return
		}
		c.JSON(http.StatusInternalServerError, Message("gagal mendapatkan user"))
		return
	}
	c.IndentedJSON(http.StatusOK, Message("berhasil login"))
}