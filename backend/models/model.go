package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Penjual struct{
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Nama string `json:"nama" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	NoHp string `json:"nohp" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type Pembeli struct{
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Nama string `json:"nama" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	NoHp string `json:"nohp" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type Barang struct{
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NamaBarang string `json:"namabarang" validate:"required"`
	Jenis string `json:"jenis" validate:"required"`
	Harga uint32 `json:"harga" validate:"required"`
	Stok int `json:"stok" validate:"required"`
	PenjualID primitive.ObjectID `json:"penjual_id" validate:"required"`
}

