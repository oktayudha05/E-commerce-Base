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

type ResBarang struct{
	ID string `json:"id" bson:"_id"`
	NamaBarang string `json:"nama_barang" bson:"namabarang"`
	Jenis string `json:"jenis" bson:"jenis"`
	Harga uint32 `json:"harga" bson:"harga"`
	Stok int `json:"stok" bson:"stok"`
	NamaPenjual string `json:"nama_penjual" bson:"nama_penjual"`
	AlamatPenjual string `json:"alamat_penjual" bson:"alamat_penjual"`
}

