package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func init(){
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("gagal memuat .env")
	}
	MONGO_URI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("gagal konek ke database")
	}
	Db = client.Database("pkm-sayur")
}