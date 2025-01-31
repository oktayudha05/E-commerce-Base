package models

type Penjual struct{
	Nama string `json:"nama" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	NoHp string `json:"nohp" validate:"required"`
}

type Pembeli struct{
	Nama string `json:"nama" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	NoHp string `json:"nohp" validate:"required"`
}
