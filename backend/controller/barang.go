package controller

import (
	"backend/models"
	"backend/utils"
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
		c.JSON(http.StatusBadRequest, utils.Message("request tidak sesuai format"))
		return
	}
	session := sessions.Default(c)
	idPenjualHex := session.Get("penjual_id")
	idPenjual, err := primitive.ObjectIDFromHex(idPenjualHex.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Message("gagal menentukan id"))
		return
	}
	req.PenjualID = idPenjual
	
	filter := bson.M{"namabarang": req.NamaBarang, "jenis": req.Jenis, "penjualid": req.PenjualID}
	tambah := bson.M{"$inc": bson.M{"stok": req.Stok}}
	result := collBarang.FindOneAndUpdate(ctx, filter, tambah)
	if result.Err() == nil {
		c.IndentedJSON(http.StatusCreated, utils.Message("berhasil tambah stok barang di database", req))
		return
	}
	req.ID = primitive.NewObjectID()
	_, err = collBarang.InsertOne(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Message("gagal menambahkan barang ke database"))
		return
	}
	c.IndentedJSON(http.StatusOK, utils.Message("berhasil tambah barang", req))
}

func GetAllBarang(c *gin.Context){
	ctx := c.Request.Context()
	
	pipeline := bson.A{
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "penjual"},
				{Key: "localField", Value: "penjualid"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "penjual"},
			}},
		},
		bson.D{{Key: "$unwind", Value: "$penjual"}},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "id", Value: 1},
				{Key: "namabarang", Value: 1},
				{Key: "jenis", Value: 1},
				{Key: "harga", Value: 1},
				{Key: "stok", Value: 1},
				{Key: "nama_penjual", Value: "$penjual.nama"},
				{Key: "alamat_penjual", Value: "$penjual.alamat"},
			}},
		},
	}

	cursor, err := collBarang.Aggregate(ctx, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Message("Gagal mengambil data barang"))
		return
	}
	defer cursor.Close(ctx)
	var results []models.ResBarang
	err = cursor.All(ctx, &results); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Message("Gagal memproses data barang"))
		return
	}
	if len(results) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": []interface{}{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": results})
}

func GetBarangById(c *gin.Context){
	ctx := c.Request.Context()
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Message("ID tidak valid"))
		return
	}
	pipeline := bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: objectId}}}},
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: "penjual"},
				{Key: "localField", Value: "penjualid"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "penjual"},
			}},
		},
		bson.D{{Key: "$unwind", Value: "$penjual"}},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "id", Value: 1},
				{Key: "namabarang", Value: 1},
				{Key: "jenis", Value: 1},
				{Key: "harga", Value: 1},
				{Key: "stok", Value: 1},
				{Key: "nama_penjual", Value: "$penjual.nama"},
				{Key: "alamat_penjual", Value: "$penjual.alamat"},
			}},
		},
	}
	cursor, err := collBarang.Aggregate(ctx, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Message("Gagal mengambil data"))
		return
	}
	defer cursor.Close(ctx)
	var result models.ResBarang
	if cursor.Next(ctx){
		err := cursor.Decode(&result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.Message("Gagal memproses data"))
			return
		}
	} else {
		c.JSON(http.StatusNotFound, utils.Message("Barang tidak ditemukan"))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}