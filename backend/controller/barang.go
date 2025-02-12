package controller

import (
	"backend/models"
	. "backend/utils"
	"net/http"

	"backend/database"

	"github.com/gin-contrib/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

var collBarang = database.Db.Collection("barang")


func AddBarang(c *gin.Context){
	ctx := c.Request.Context()
	var req models.Barang
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Message("request tidak sesuai format"))
		return
	}
	session := sessions.Default(c)
	idPenjualHex := session.Get("penjual_id")
	idPenjual, err := primitive.ObjectIDFromHex(idPenjualHex.(string))
	req.PenjualID = idPenjual
	
	filter := bson.M{"namabarang": req.NamaBarang, "jenis": req.Jenis, "penjualid": req.PenjualID}
	tambah := bson.M{"$inc": bson.M{"stok": req.Stok}}
	result := collBarang.FindOneAndUpdate(ctx, filter, tambah)
	if result.Err() == nil {
		c.IndentedJSON(http.StatusCreated, Message("berhasil tambah stok barang di database", req))
		return
	}
	req.ID = primitive.NewObjectID()
	_, err = collBarang.InsertOne(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Message("gagal menambahkan barang ke database"))
		return
	}
	c.IndentedJSON(http.StatusOK, Message("berhasil tambah barang", req))
}