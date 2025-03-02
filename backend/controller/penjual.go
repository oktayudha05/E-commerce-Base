package controller

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
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
		c.JSON(http.StatusConflict, utils.Message("gagal bind data"))
		return
	}
	err = validate.Struct(penjual)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Message("format data salah"))
		return
	}
	count, err := collPenjual.CountDocuments(ctx, bson.M{"username": penjual.Username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Message("error mencari username"))
		return
	}
	if count > 0 {
		c.JSON(http.StatusFound, utils.Message("Username sudah ada"))
		return
	}
	hashPass, err := utils.HashPassword(penjual.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Message("gagal hash password"))
		return
	}
	penjual.Password = hashPass	
	_, err = collPenjual.InsertOne(ctx, penjual)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Message("gagal tambah user", gin.H{"data": err}))
		return
	}

	c.IndentedJSON(http.StatusOK, utils.Message("berhasil nambah penjual", penjual))
}

func LoginPenjual(c *gin.Context){
	ctx := c.Request.Context()
	var reqPenjual struct{
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	err := c.BindJSON(&reqPenjual)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Message("gagal bind request"))
		return
	}
	err = validate.Struct(reqPenjual)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Message("data tidak sesuai format"))
		return
	}
	var penjual models.Penjual
	filter := bson.M{"username": reqPenjual.Username}
	err = collPenjual.FindOne(ctx, filter).Decode(&penjual)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, utils.Message("username tidak ada"))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.Message("gagal mendapatkan user"))
		return
	}
	err = utils.CekPassword(penjual.Password, reqPenjual.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.Message("password salah"))
		return
	}

	session := sessions.Default(c)
	session.Set("username", penjual.Username)
	session.Set("penjual_id", penjual.ID.Hex())
	session.Set("role", "penjual")
	session.Save()

	c.IndentedJSON(http.StatusOK, utils.Message("berhasil login"))
}